package app

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-svc-tpl/app/controller"
)

func addRoutes() {
	api := e.Group("api")

	api.GET("/doc/*", echoSwagger.WrapHandler)

	api.GET("/foo", controller.Foo)
}
