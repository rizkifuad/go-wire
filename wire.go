//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rizkix/wired/config"
	"github.com/rizkix/wired/controller"
	"github.com/rizkix/wired/delivery/grpc"
	"github.com/rizkix/wired/delivery/http"
	"github.com/rizkix/wired/plugin"
	"github.com/rizkix/wired/repo"
)

func InitializeApp() (App, error) {
	wire.Build(NewApp, config.New, controller.New, repo.New, plugin.NewMysqlConnection, http.New, grpc.New)
	return App{}, nil
}
