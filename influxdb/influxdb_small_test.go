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

	"github.com/intelsdi-x/snap-plugin-collector-influxdb/influxdb/dtype"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

type mcMock struct {
	mock.Mock
}

func (mc *mcMock) GetStatistics() (dtype.Results, error) {
	args := mc.Called()
	return args.Get(0).(dtype.Results), args.Error(1)
}

func (mc *mcMock) GetDiagnostics() (dtype.Results, error) {
	args := mc.Called()
	return args.Get(0).(dtype.Results), args.Error(1)
}

func (mc *mcMock) InitURLs(string, int64, string, string) error {
	args := mc.Called()
	return args.Error(0)
}

var mockStats = dtype.Results{

	"shard/1": &dtype.Series{
		Data: map[string]interface{}{
			"columnA": 1,
			"columnB": 10.1,
			"columnC": "value",
		},
		Tags: map[string]string{
			"tag1": "v1",
			"tag2": "v2",
		},
	},
	"httpd": &dtype.Series{
		Data: map[string]interface{}{
			"columnA": 1,
			"columnB": 10.1,
			"columnC": "value",
		},
		Tags: map[string]string{
			"tag1": "v1",
			"tag2": "v2",
		},
	},
}

var mockDiagn = dtype.Results{
	"build": &dtype.Series{
		Data: map[string]interface{}{
			"columnA": 1,
			"columnB": 10.1,
			"columnC": "value",
		},
	},
}

var mockMtsStat = []plugin.Metric{
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "stat", "shard", "1", "columnA")},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "stat", "shard", "1", "columnB")},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "stat", "shard", "1", "columnC")},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "stat", "httpd", "columnA")},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "stat", "httpd", "columnB")},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "stat", "httpd", "columnC")},
}

var mockMtsDiagn = []plugin.Metric{
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "diagn", "build", "columnA")},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "diagn", "build", "columnB")},
	plugin.Metric{Namespace: plugin.NewNamespace("intel", "influxdb", "diagn", "build", "columnC")},
}

// Mts is a mocked metrics, both statistical and diagnostic
var mockMts = append(mockMtsStat, mockMtsDiagn...)

func TestGetConfigPolicy(t *testing.T) {
	influxdbPlugin := New()

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
			mc := &mcMock{}
			influxdbPlugin := &InfluxdbCollector{initialized: false, service: mc, data: map[string]datum{}}
			cfg := getMockConfig()

			mc.On("InitURLs").Return(errors.New("x"))

			So(func() { influxdbPlugin.GetMetricTypes(cfg) }, ShouldNotPanic)
			results, err := influxdbPlugin.GetMetricTypes(cfg)
			So(err, ShouldNotBeNil)
			So(results, ShouldBeEmpty)
		})
	})

	Convey("Metrics are not available", t, func() {
		Convey("when cannot obtain any data", func() {
			mc := &mcMock{}
			influxdbPlugin := &InfluxdbCollector{initialized: false, service: mc, data: map[string]datum{}}
			cfg := getMockConfig()

			mc.On("InitURLs").Return(nil)
			mc.On("GetStatistics").Return(dtype.Results{}, errors.New("x"))
			mc.On("GetDiagnostics").Return(dtype.Results{}, errors.New("x"))

			So(func() { influxdbPlugin.GetMetricTypes(cfg) }, ShouldNotPanic)
			results, err := influxdbPlugin.GetMetricTypes(cfg)
			So(err, ShouldNotBeNil)
			So(results, ShouldBeEmpty)
		})
		Convey("when cannot obtain statistics data", func() {
			mc := &mcMock{}
			influxdbPlugin := &InfluxdbCollector{initialized: false, service: mc, data: map[string]datum{}}
			cfg := getMockConfig()

			mc.On("InitURLs").Return(nil)
			mc.On("GetStatistics").Return(dtype.Results{}, errors.New("x"))
			mc.On("GetDiagnostics").Return(mockDiagn, nil)

			So(func() { influxdbPlugin.GetMetricTypes(cfg) }, ShouldNotPanic)
			results, err := influxdbPlugin.GetMetricTypes(cfg)
			So(err, ShouldNotBeNil)
			So(results, ShouldBeEmpty)
		})
		Convey("when cannot obtain diagnostics data", func() {
			mc := &mcMock{}
			influxdbPlugin := &InfluxdbCollector{initialized: false, service: mc, data: map[string]datum{}}
			cfg := getMockConfig()

			mc.On("InitURLs").Return(nil)
			mc.On("GetStatistics").Return(mockStats, nil)
			mc.On("GetDiagnostics").Return(dtype.Results{}, errors.New("x"))

			So(func() { influxdbPlugin.GetMetricTypes(cfg) }, ShouldNotPanic)
			results, err := influxdbPlugin.GetMetricTypes(cfg)
			So(err, ShouldNotBeNil)
			So(results, ShouldBeEmpty)
		})
	})
	Convey("Successful getting metrics types", t, func() {
		mc := &mcMock{}
		influxdbPlugin := &InfluxdbCollector{initialized: false, service: mc, data: map[string]datum{}}
		cfg := getMockConfig()

		mc.On("InitURLs").Return(nil)
		mc.On("GetStatistics").Return(mockStats, nil)
		mc.On("GetDiagnostics").Return(mockDiagn, nil)

		So(func() { influxdbPlugin.GetMetricTypes(cfg) }, ShouldNotPanic)
		results, err := influxdbPlugin.GetMetricTypes(cfg)
		So(err, ShouldBeNil)
		So(results, ShouldNotBeEmpty)
	})
}

func TestCollectMetrics(t *testing.T) {

	mts := getMockMetricsConfigured()

	Convey("Initialization fails", t, func() {
		mc := &mcMock{}
		influxdbPlugin := &InfluxdbCollector{initialized: false, service: mc, data: map[string]datum{}}

		mc.On("InitURLs").Return(errors.New("x"))

		So(func() { influxdbPlugin.CollectMetrics(mts) }, ShouldNotPanic)
		results, err := influxdbPlugin.CollectMetrics(mts)
		So(err, ShouldNotBeNil)
		So(results, ShouldBeEmpty)
	})
	Convey("Metrics are not available", t, func() {
		mc := &mcMock{}
		influxdbPlugin := &InfluxdbCollector{initialized: false, service: mc, data: map[string]datum{}}
		mc.On("InitURLs").Return(nil)

		Convey("when cannot get statistics data", func() {
			mc.On("GetStatistics").Return(dtype.Results{}, errors.New("x"))
			mc.On("GetDiagnostics").Return(mockDiagn, nil)

			results, err := influxdbPlugin.CollectMetrics(mts)
			So(err, ShouldNotBeNil)
			So(results, ShouldBeEmpty)
		})
		Convey("when cannot get diagnostics data", func() {
			influxdbPlugin.initialized = false
			mc.On("GetStatistics").Return(mockStats, nil)
			mc.On("GetDiagnostics").Return(dtype.Results{}, errors.New("x"))

			results, err := influxdbPlugin.CollectMetrics(mts)
			So(err, ShouldNotBeNil)
			So(results, ShouldBeEmpty)
		})
	})
	Convey("Successful collecting metrics", t, func() {
		mc := &mcMock{}
		influxdbPlugin := &InfluxdbCollector{initialized: false, service: mc, data: map[string]datum{}}
		mc.On("InitURLs").Return(nil)
		mc.On("GetStatistics").Return(mockStats, nil)
		mc.On("GetDiagnostics").Return(mockDiagn, nil)

		results, err := influxdbPlugin.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(results, ShouldNotBeEmpty)
		So(len(results), ShouldEqual, len(mockMts))
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

func getMockMetricsConfigured() []plugin.Metric {
	mts := mockMts
	cfg := getMockConfig()

	// add mocked config to each metric
	for i := range mts {
		mts[i].Config = cfg
	}

	return mts
}
