package controller

import (
	"github.com/labstack/echo/v4"
	"go-svc-tpl/model"
	"net/http"
)

// @tags Foo
// @summary A demo API
// @router /foo [get]
// @produce json
func foo(c echo.Context) error {
	// just a demo
	var result model.Foo
	model.DB.Model(&model.Foo{}).First(&result)

	return c.JSON(http.StatusOK, &struct {
		Message string `json:"message"`
	}{Message: "Hello World! 你好世界！"})

}
