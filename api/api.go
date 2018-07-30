package api

import (
	"github.com/gedelumbung/go-catalog/config"
	"github.com/gedelumbung/go-catalog/repository"
	"github.com/labstack/echo"
)

type API struct {
	config *conf.Configuration
	web    *echo.Echo
	db     repository.Repository
}

func (a *API) ListenAndServe() {
	a.web.Logger.Fatal(a.web.Start(a.config.API.Host))
}

func NewAPI(config *conf.Configuration, db repository.Repository) *API {
	a := &API{config: config, web: echo.New(), db: db}
	a.registerRoutes()
	return a
}
