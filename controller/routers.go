package controller

import (
	echoSwagger "github.com/swaggo/echo-swagger"
)

func addRoutes() {
	api := e.Group("api")

	api.GET("/doc/*", echoSwagger.WrapHandler)

	api.GET("/foo", foo)
}
