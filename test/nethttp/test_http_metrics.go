// Copyright (c) 2024 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/alibaba/opentelemetry-go-auto-instrumentation/pkg/verifier"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"net/http"
	"strconv"
	"time"
)

func setupMetricHttp() {
	http.HandleFunc("/a", helloHandler)
	var err error
	port, err = verifier.GetFreePort()
	if err != nil {
		panic(err)
	}
	err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	go setupMetricHttp()
	time.Sleep(1 * time.Second)
	_, err := http.Get("http://127.0.0.1:" + strconv.Itoa(port) + "/a")
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)
	verifier.WaitAndAssertMetrics("http.server.request.duration", func(mrs metricdata.ResourceMetrics) {
		if len(mrs.ScopeMetrics) <= 0 {
			panic("No http.server.request.duration metrics received!")
		}
		point := mrs.ScopeMetrics[0].Metrics[0].Data.(metricdata.Histogram[float64])
		if point.DataPoints[0].Count != 1 {
			panic("http.server.request.duration metrics count is not 1")
		}
		verifier.VerifyHttpServerMetricsAttributes(point.DataPoints[0].Attributes.ToSlice(), "GET", "/a", "", "http", "1.1", "http", 200)
	})

	verifier.WaitAndAssertMetrics("http.client.request.duration", func(mrs metricdata.ResourceMetrics) {
		if len(mrs.ScopeMetrics) <= 0 {
			panic("No http.server.request.duration metrics received!")
		}
		point := mrs.ScopeMetrics[0].Metrics[0].Data.(metricdata.Histogram[float64])
		if point.DataPoints[0].Count != 1 {
			panic("http.client.request.duration metrics count is not 1")
		}
		verifier.VerifyHttpClientMetricsAttributes(point.DataPoints[0].Attributes.ToSlice(), "GET", "127.0.0.1:"+strconv.Itoa(port), "", "http", "1.1", port, 200)
	})
}
