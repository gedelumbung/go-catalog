package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (a *API) registerRoutes() {
	a.web.Pre(middleware.RemoveTrailingSlash())
	a.web.Use(middleware.CORS())
	a.web.Use(middleware.Recover())

	g := a.web.Group("/v1")
	g.GET("/ping", func(c echo.Context) error {
		type Hello struct {
			Ping    string `json:"ping"`
			AppName string `json:"app_name"`
			Version string `json:"version"`
		}
		hello := &Hello{
			Ping:    "pong",
			AppName: "Catalog",
			Version: "1.0",
		}
		return c.JSON(http.StatusOK, hello)
	})

	product := g.Group("/products")
	product.GET("", a.GetAllProducts)
	product.GET("/:id", a.GetProduct)
	product.GET("/:id/images/:image_id", a.GetImage)
	product.POST("", a.StoreProduct)
	product.DELETE("/:id", a.DeleteProduct)
}
