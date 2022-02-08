// @title        Golang Service Template
// @version      0.1
// @description  Golang back-end service template, get started with back-end projects quickly
// @BasePath     /api

package main

import (
	"go-svc-tpl/app"
	_ "go-svc-tpl/docs"
	"go-svc-tpl/model"
)

func main() {
	model.Init()
	app.InitWebFramework()
	app.StartServer()
}
