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

package echo

import (
	"github.com/alibaba/opentelemetry-go-auto-instrumentation/pkg/inst-api/instrumenter"

	"github.com/alibaba/opentelemetry-go-auto-instrumentation/pkg/api"
	echo "github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/sdk/trace"
)

var echoEnabler = instrumenter.NewDefaultInstrumentEnabler()

func otelTraceMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if err = next(c); err != nil {
				c.Error(err)
			}
			lcs := trace.LocalRootSpanFromGLS()
			if lcs != nil && c.Path() != "" && c.Request() != nil && c.Request().URL != nil && (c.Request().URL.Path != c.Path()) {
				lcs.SetName(c.Path())
			}
			return
		}
	}
}

func afterNewEcho(call api.CallContext, e *echo.Echo) {
	if !echoEnabler.Enable() {
		return
	}
	if e == nil {
		return
	}

	e.Use(otelTraceMiddleware())
}
