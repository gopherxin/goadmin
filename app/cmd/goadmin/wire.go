//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/gopherxin/goadmin/app/internal/data"
	"github.com/gopherxin/goadmin/app/internal/server"
	"github.com/gopherxin/goadmin/app/internal/service"
	"goadmin/app/internal/biz"
	"goadmin/app/internal/conf"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
