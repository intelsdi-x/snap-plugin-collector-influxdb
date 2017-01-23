// +build small

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
	"testing"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	"strings"

	"net/url"

	. "github.com/smartystreets/goconvey/convey"
)

func getMockHTTPResponse(url string) ([]byte, error) {
	if strings.Contains(url, "stats") {
		return []byte(mockStatResults), nil
	} else if strings.Contains(url, "diagnostics") {
		return []byte(mockDiagnosticResults), nil
	}
	return nil, errors.New("invalid arg")
}

func getEmptyMockHTTPResponse(url string) ([]byte, error) {
	if strings.Contains(url, "stats") {
		return []byte("{}"), nil
	} else if strings.Contains(url, "diagnostics") {
		return []byte("{}"), nil
	}
	return nil, errors.New("invalid arg")
}

var mockMtsStat = []plugin.Metric{
	plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "influxdb", "stat", "shard", "diskBytes"),
		Tags:      map[string]string{},
	},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "stat", "shard", "fieldsCreate"),
		Tags: map[string]string{},
	},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "stat", "shard", "seriesCreate"),
		Tags: map[string]string{},
	},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "stat", "httpd", "queryReq"),
		Tags: map[string]string{},
	},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "stat", "httpd", "req"),
		Tags: map[string]string{}},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "stat", "httpd", "reqActive"),
		Tags: map[string]string{},
	},
}

var mockMtsDiagn = []plugin.Metric{
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "diagn", "build", "Branch"),
		Tags: map[string]string{},
	},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "diagn", "build", "Commit"),
		Tags: map[string]string{},
	},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "diagn", "build", "Version"),
		Tags: map[string]string{},
	},
}

// Mts is a mocked metrics, both statistical and diagnostic
var mockMts = append(mockMtsStat, mockMtsDiagn...)

func TestGetConfigPolicy(t *testing.T) {
	influxdbPlugin := &influxdbCollector{
		getResponse: getMockHTTPResponse,
	}

	Convey("Getting config policy", t, func() {
		So(func() { influxdbPlugin.GetConfigPolicy() }, ShouldNotPanic)
		policy, err := influxdbPlugin.GetConfigPolicy()
		So(err, ShouldBeNil)
		So(policy, ShouldNotBeNil)
	})
}

func TestGetMetricTypes(t *testing.T) {

	Convey("Initialization fails", t, func() {
		Convey("when no config items available", func() {
			influxdbPlugin := New()
			cfg := plugin.Config{}

			So(func() { influxdbPlugin.GetMetricTypes(cfg) }, ShouldNotPanic)
			results, err := influxdbPlugin.GetMetricTypes(cfg)
			So(err, ShouldNotBeNil)
			So(results, ShouldBeEmpty)
		})
		Convey("when one of config item is not available", func() {
			influxdbPlugin := New()
			cfg := getMockConfig()
			delete(cfg, "user")

			So(func() { influxdbPlugin.GetMetricTypes(cfg) }, ShouldNotPanic)
			results, err := influxdbPlugin.GetMetricTypes(cfg)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "config item not found")
			So(results, ShouldBeEmpty)
		})
		Convey("when config item has different type than expected", func() {
			influxdbPlugin := New()
			cfg := getMockConfig()
			// set a valid value as a port (expected int64 type)
			cfg["port"] = "1234"

			So(func() { influxdbPlugin.GetMetricTypes(cfg) }, ShouldNotPanic)
			results, err := influxdbPlugin.GetMetricTypes(cfg)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "config item is not an int64")
			So(results, ShouldBeEmpty)
		})
		Convey("when initialization of URLs returns error", func() {
			influxdbPlugin := New()
			cfg := getMockConfig()

			So(func() { influxdbPlugin.GetMetricTypes(cfg) }, ShouldNotPanic)
			results, err := influxdbPlugin.GetMetricTypes(cfg)
			So(err, ShouldNotBeNil)
			So(results, ShouldBeEmpty)
		})
	})

	Convey("Metrics are not available", t, func() {
		Convey("when cannot obtain any data", func() {
			influxdbPlugin := &influxdbCollector{
				getResponse:   getMockHTTPResponse,
				urlDiagnostic: &url.URL{Path: ""},
				urlStatistic:  &url.URL{Path: ""},
			}
			cfg := getMockConfig()

			So(func() { influxdbPlugin.GetMetricTypes(cfg) }, ShouldNotPanic)
			results, err := influxdbPlugin.GetMetricTypes(cfg)
			So(results, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})
	})
	Convey("Successfully get metrics types", t, func() {
		influxdbPlugin := &influxdbCollector{
			getResponse:   getMockHTTPResponse,
			urlDiagnostic: &url.URL{Path: "diagnostics"},
			urlStatistic:  &url.URL{Path: "stats"},
		}
		cfg := getMockConfig()

		So(func() { influxdbPlugin.GetMetricTypes(cfg) }, ShouldNotPanic)
		results, err := influxdbPlugin.GetMetricTypes(cfg)
		So(err, ShouldBeNil)
		So(results, ShouldNotBeEmpty)
	})
}

func TestCollectMetrics(t *testing.T) {

	Convey("Initialization fails", t, func() {
		influxdbPlugin := &influxdbCollector{
			getResponse:   getMockHTTPResponse,
			urlDiagnostic: &url.URL{Path: ""},
			urlStatistic:  &url.URL{Path: ""},
		}

		So(func() { influxdbPlugin.CollectMetrics(mockMts) }, ShouldNotPanic)
		results, err := influxdbPlugin.CollectMetrics(mockMts)
		So(err, ShouldNotBeNil)
		So(results, ShouldBeEmpty)
	})

	Convey("Metrics are not available", t, func() {
		influxdbPlugin := &influxdbCollector{
			getResponse:   getEmptyMockHTTPResponse,
			urlDiagnostic: &url.URL{Path: "diagnostics"},
			urlStatistic:  &url.URL{Path: "stats"},
		}

		Convey("when cannot get  data", func() {
			results, err := influxdbPlugin.CollectMetrics(mockMts)
			So(err, ShouldBeNil)
			So(results, ShouldBeEmpty)
		})
	})
	Convey("Successful collecting metrics", t, func() {
		influxdbPlugin := &influxdbCollector{
			getResponse:   getMockHTTPResponse,
			urlDiagnostic: &url.URL{Path: "diagnostics"},
			urlStatistic:  &url.URL{Path: "stats"},
		}

		results, err := influxdbPlugin.CollectMetrics(mockMts)
		So(err, ShouldBeNil)
		So(results, ShouldNotBeEmpty)
		for _, i := range results {
			for _, y := range i.Namespace.Strings() {
				print(y)
				print("/")
			}
			println("")
		}
		So(len(results), ShouldBeGreaterThanOrEqualTo, len(mockMts))
	})
}

func getMockConfig() plugin.Config {
	// mocking config
	return plugin.Config{
		"host":     "hostname",
		"port":     int64(1234),
		"user":     "test",
		"password": "passwd",
	}
}
