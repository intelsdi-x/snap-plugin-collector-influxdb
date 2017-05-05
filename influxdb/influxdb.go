/*
http://www.apache.org/licenses/LICENSE-2.0.txt
Copyright 2016 Intel Corporation
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package influxdb

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"

	log "github.com/Sirupsen/logrus"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	// Name of plugin
	Name = "influxdb"
	// Version of plugin
	Version = 7

	nsVendor = "intel"
	nsClass  = "influxdb"

	nsTypeStats = "stat"
	nsTypeDiagn = "diagn"
)

const (
	typeUnknown = iota
	typeStats
	typeDiagn
)

// prefix in metric namespace
var prefix = []string{nsVendor, nsClass}

type getResponse func(url string) ([]byte, error)

// influxdbCollector holds data retrieved from influxDB system monitoring
type influxdbCollector struct {
	urlStatistic  *url.URL
	urlDiagnostic *url.URL
	getResponse
}

// New returns new instance of snap-plugin-collector-influxdb
func New() plugin.Collector {
	return &influxdbCollector{
		getResponse: getHttpResponse,
	}
}

// GetConfigPolicy returns a ConfigPolicy
func (ic *influxdbCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	cfgKey := []string{"intel", "influxdb"}
	policy.AddNewStringRule(cfgKey, "host", false, plugin.SetDefaultString("localhost"))
	policy.AddNewIntRule(cfgKey, "port", false, plugin.SetDefaultInt(8086))
	policy.AddNewStringRule(cfgKey, "user", false, plugin.SetDefaultString("admin"))
	policy.AddNewStringRule(cfgKey, "password", false, plugin.SetDefaultString("admin"))
	return *policy, nil
}

// GetMetricTypes returns list of metrics based on influxDB system monitoring
func (ic *influxdbCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	if ic.urlDiagnostic == nil || ic.urlStatistic == nil {
		if err := ic.init(cfg); err != nil {
			return nil, err
		}
	}

	return ic.getMetrics()
}

// CollectMetrics collects given metrics
func (ic *influxdbCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	res := []plugin.Metric{}
	if len(mts) == 0 {
		return nil, errors.New("No metrics requested")
	}
	if ic.urlDiagnostic == nil || ic.urlStatistic == nil {
		ic.init(mts[0].Config)
	}

	metrics, err := ic.getMetrics()
	if err != nil {
		return nil, err
	}

	// return only requested metrics
	ts := time.Now()
	for _, req := range mts {
		for _, metric := range metrics {
			if reflect.DeepEqual(req.Namespace.Strings(), metric.Namespace.Strings()) {
				// merge any new tags
				for k, v := range metric.Tags {
					req.Tags[k] = v
				}
				req.Data = metric.Data
				req.Timestamp = ts
				res = append(res, req)
			}
		}
	}

	return res, nil
}

// init initializes InfluxdbCollector instance based on plugin config `cfg`
func (ic *influxdbCollector) init(cfg plugin.Config) error {
	host, err := cfg.GetString("host")
	if err != nil {
		return fmt.Errorf("Cannot get a hostname from plugin config, err=%s", err.Error())
	}

	port, err := cfg.GetInt("port")
	if err != nil {
		return fmt.Errorf("Cannot get a port from plugin config, err=%s", err.Error())
	}

	user, err := cfg.GetString("user")
	if err != nil {
		return fmt.Errorf("Cannot get a username from plugin config, err=%s", err.Error())
	}

	passwd, err := cfg.GetString("password")
	if err != nil {
		return fmt.Errorf("Cannot get a password from plugin config, err=%s", err.Error())
	}

	if err := ic.InitURLs(host, port, user, passwd); err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"function": "init",
	}).Info("Succeeded plugin initialization")
	return nil
}

func (ic *influxdbCollector) getMetrics() ([]plugin.Metric, error) {
	stats, err := ic.getStatistics()
	if err != nil {
		return nil, err
	}
	diags, err := ic.getDiagnostics()
	if err != nil {
		return nil, err
	}

	return append(diags, stats...), nil
}

// getDiagnostics executes the command "SHOW DIAGNOSTICS" (indirectly)
func (ic *influxdbCollector) getDiagnostics() ([]plugin.Metric, error) {
	mts := []plugin.Metric{}
	var diag diagnostics
	response, err := ic.getResponse(ic.urlDiagnostic.String())
	if err != nil {
		log.Errorf("error getting response err=%v response=%v", err.Error(),
			response)
		return nil, err
	}
	err = json.Unmarshal(response, &diag)
	if err != nil {
		return nil, err
	}
	for _, result := range diag.Results {
		for _, series := range result.Series {
			for _, values := range series.Values {
				for idx, value := range values {
					mts = append(mts, plugin.Metric{
						Namespace: plugin.NewNamespace(nsVendor, nsClass,
							nsTypeDiagn, series.Name, series.Columns[idx]),
						Data: value,
					})
				}
			}
		}
	}
	return mts, nil
}

// getStatistics executes the command "SHOW STATS" (indirectly)
func (ic *influxdbCollector) getStatistics() ([]plugin.Metric, error) {
	mts := []plugin.Metric{}
	var stats stats
	response, err := ic.getResponse(ic.urlStatistic.String())
	if err != nil {
		log.Errorf("error getting response err=%v response=%v", err.Error(),
			response)
		return nil, err
	}
	err = json.Unmarshal(response, &stats)
	if err != nil {
		return nil, err
	}
	for _, result := range stats.Results {
		for _, series := range result.Series {
			for _, values := range series.Values {
				for idx, value := range values {
					mts = append(mts, plugin.Metric{
						Namespace: plugin.NewNamespace(nsVendor, nsClass,
							nsTypeStats, series.Name, series.Columns[idx]),
						Data: value,
						Tags: series.Tags,
					})
				}
			}
		}
	}
	return mts, nil
}

// InitURLs initializes URLs based on settings
func (ic *influxdbCollector) InitURLs(host string, port int64, user string, passwd string) error {
	errs := []error{}
	var err error
	queryStatementStats := "show stats"
	queryStatementDiagn := "show diagnostics"

	if ic.urlStatistic, err = createURL(host, port, user, passwd, queryStatementStats); err != nil {
		errs = append(errs, err)

		log.WithFields(log.Fields{
			"block":    "monitor",
			"function": "InitURLs",
			"err":      err,
		}).Errorf("Cannot parse raw url into a URL structure with query `%s`", queryStatementStats)
	}

	if ic.urlDiagnostic, err = createURL(host, port, user, passwd, queryStatementDiagn); err != nil {
		errs = append(errs, err)

		log.WithFields(log.Fields{
			"block":    "monitor",
			"function": "InitURLs",
			"err":      err,
		}).Errorf("Cannot parse raw url into a URL structure with query `%s`", queryStatementDiagn)
	}

	if len(errs) != 0 {
		return errors.New("Cannot initialize URLs, invalid URL-encoding")
	}

	return nil
}

// --------- helper functions -------------- //

func getHttpResponse(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

// createURL returns URL structure created base on hostname, port, credentials and query statement
func createURL(host string, port int64, user string, passwd string, query string) (*url.URL, error) {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d/query?u=%s&p=%s&pretty=true",
		host,
		port,
		user,
		passwd,
	))

	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("q", query)
	u.RawQuery = q.Encode()

	return u, nil
}
