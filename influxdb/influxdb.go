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
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	"github.com/intelsdi-x/snap-plugin-collector-influxdb/influxdb/dtype"
	"github.com/intelsdi-x/snap-plugin-collector-influxdb/influxdb/monitor"
)

const (
	// Name of plugin
	Name = "influxdb"
	// Version of plugin
	Version = 5

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

// InfluxdbCollector holds data retrieved from influxDB system monitoring
type InfluxdbCollector struct {
	data        map[string]datum
	service     monitor.Monitoring
	initialized bool
}

type datum struct {
	value interface{}
	tags  map[string]string
}

// New returns new instance of snap-plugin-collector-influxdb
func New() plugin.Collector {
	return &InfluxdbCollector{initialized: false, service: &monitor.Monitor{}, data: map[string]datum{}}
}

// GetConfigPolicy returns a ConfigPolicy
func (ic *InfluxdbCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	cfgKey := []string{"intel", "influxdb"}
	policy.AddNewStringRule(cfgKey, "host", false, plugin.SetDefaultString("localhost"))
	policy.AddNewIntRule(cfgKey, "port", false, plugin.SetDefaultInt(8086))
	policy.AddNewStringRule(cfgKey, "user", true)
	policy.AddNewStringRule(cfgKey, "password", true)
	return *policy, nil
}

// GetMetricTypes returns list of metrics based on influxDB system monitoring
func (ic *InfluxdbCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	mts := []plugin.Metric{}
	if err := ic.init(cfg); err != nil {
		return nil, err
	}

	// get InfluxDB internal statistics
	if err := ic.getStatistics(); err != nil {
		return nil, fmt.Errorf("Cannot get influxdb internal statistics, err=%s", err.Error())
	}

	// get InfluxDB internal diagnostics info
	if err := ic.getDiagnostics(); err != nil {
		return nil, fmt.Errorf("Cannot get influxdb diagnostic information, err=%s", err.Error())
	}

	for key, dat := range ic.data {
		mts = append(mts, plugin.Metric{
			Namespace: plugin.NewNamespace(prefix...).AddStaticElements(splitKey(key)...),
			Tags:      dat.tags,
			Version:   Version,
		})
	}

	return mts, nil
}

// CollectMetrics collects given metrics
func (ic *InfluxdbCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	if !ic.initialized {
		// mts has one item at least if CollectMetrics() has been called
		if err := ic.init(mts[0].Config); err != nil {
			return nil, err
		}
		// get diagnostic information (once only)
		if err := ic.getDiagnostics(); err != nil {
			return nil, fmt.Errorf("Cannot get influxdb diagnostic information, err=%s", err.Error())
		}
	}

	// get statistics
	if err := ic.getStatistics(); err != nil {
		return nil, fmt.Errorf("Cannot get influxdb internal statistics, err=%s", err.Error())
	}

	for i := range mts {
		ns := mts[i].Namespace
		if dat, ok := ic.data[reflectKey(ns)]; ok {
			mts[i].Data = dat.value
			mts[i].Timestamp = time.Now()
			mts[i].Tags = dat.tags
		} else {
			// only log about it
			log.WithFields(log.Fields{
				"function": "CollectMetrics",
				"metric":   ns,
			}).Error("No data found")
		}
	}

	return mts, nil
}

// init initializes InfluxdbCollector instance based on plugin config `cfg`
func (ic *InfluxdbCollector) init(cfg plugin.Config) error {
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

	if err := ic.service.InitURLs(host, port, user, passwd); err != nil {
		return err
	}

	ic.initialized = true
	log.WithFields(log.Fields{
		"function": "init",
	}).Info("Succeeded plugin initialization")
	return nil
}

// getDiagnostics executes the command "SHOW DIAGNOSTICS" (indirectly)
func (ic *InfluxdbCollector) getDiagnostics() error {
	return ic.getData(typeDiagn)
}

// getStatistics executes the command "SHOW STATS" (indirectly)
func (ic *InfluxdbCollector) getStatistics() error {
	return ic.getData(typeStats)
}

// getData executes a command specified by given `type` of desired data
// and assigns its results to InfluxdbCollector structure item `data`
func (ic *InfluxdbCollector) getData(kind int) error {
	var results dtype.Results
	var err error
	var nsType string

	switch kind {
	case typeStats:
		nsType = nsTypeStats
		results, err = ic.service.GetStatistics()

	case typeDiagn:
		nsType = nsTypeDiagn
		results, err = ic.service.GetDiagnostics()

	default:
		err = errors.New("Invalid type of monitoring service")
	}

	if err != nil {
		return err
	}

	for seriesName, series := range results {
		for columnName := range series.Data {
			key := createKey(nsType, seriesName, columnName)
			ic.data[key] = datum{
				value: series.Data[columnName],
				tags:  series.Tags,
			}

		}
	}

	return nil
}

// createKey returns a key which identify metric's key which is composed from metric's type (might equal `stats` or `diagn`)
// and component name; all elements are joined to a single string
func createKey(nsType, seriesName, columnName string) string {
	seriesName = strings.Replace(seriesName, "/", "_", -1)
	columnName = strings.Replace(columnName, "/", "_", -1)
	return strings.Join([]string{nsType, seriesName, columnName}, "/")
}

// reflectKey returns corresponding metric's key based on metric's namespace
func reflectKey(ns plugin.Namespace) string {
	// skip metric's prefix and join the rest of elements
	return strings.Join(ns.Strings()[len(prefix):], "/")
}

// splitKey returns a slice of the substrings between slash separator
func splitKey(key string) []string {
	return strings.Split(key, "/")
}
