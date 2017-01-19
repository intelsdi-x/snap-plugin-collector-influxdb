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

package monitor

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"

	"github.com/intelsdi-x/snap-plugin-collector-influxdb/influxdb/dtype"
	"github.com/intelsdi-x/snap-plugin-collector-influxdb/influxdb/parser"
)

const (
	queryStatementStats = "show stats"
	queryStatementDiagn = "show diagnostics"
)

// Monitoring is an interface represents data monitoring service
// (needed for mocking purposes)
type Monitoring interface {
	GetStatistics() (dtype.Results, error)
	GetDiagnostics() (dtype.Results, error)
	InitURLs(string, int64, string, string) error
}

// Monitor holds urls
type Monitor struct {
	urlStatistic  *url.URL
	urlDiagnostic *url.URL
}

// GetStatistics returns statistics information (url contains query "SHOW STATS")
func (m *Monitor) GetStatistics() (dtype.Results, error) {
	return getURLResults(m.urlStatistic.String())
}

// GetDiagnostics returns diagnostics information (url contains query "SHOW DIAGNOSTICS")
func (m *Monitor) GetDiagnostics() (dtype.Results, error) {
	return getURLResults(m.urlDiagnostic.String())
}

// InitURLs initializes URLs based on settings
func (m *Monitor) InitURLs(host string, port int64, user string, passwd string) error {
	errs := []error{}
	var err error

	if m.urlStatistic, err = createURL(host, port, user, passwd, queryStatementStats); err != nil {
		errs = append(errs, err)

		log.WithFields(log.Fields{
			"block":    "monitor",
			"function": "InitURLs",
			"err":      err,
		}).Errorf("Cannot parse raw url into a URL structure with query `%s`", queryStatementStats)
	}

	if m.urlDiagnostic, err = createURL(host, port, user, passwd, queryStatementDiagn); err != nil {
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

// getURLResults returns result of GET to the specified url which has been parsed into a dtype.Results structure
func getURLResults(url string) (dtype.Results, error) {
	response, err := getHTTPResponse(url)
	if err != nil {
		return nil, err
	}
	results, err := parser.FromJSON(response)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// getHTTPResponse returns HTTP response of GET to the specified url
func getHTTPResponse(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
