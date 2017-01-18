package main

import (
	"godb/service"
	"net/http"

	"github.com/labstack/echo"
)

func RegisterApi() {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello db")
	})

	api := e.Group("/api")
	v1 := api.Group("/v1")

	customer := v1.Group("/customer")
	customer.GET("/:id", service.Retrive)
	customer.POST("", service.Create)

}
