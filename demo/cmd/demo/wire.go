// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/SmartNMS/smart-service/demo/internal/biz"
	"github.com/SmartNMS/smart-service/demo/internal/conf"
	"github.com/SmartNMS/smart-service/demo/internal/data"
	"github.com/SmartNMS/smart-service/demo/internal/server"
	"github.com/SmartNMS/smart-service/demo/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, trace.TracerProvider, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
