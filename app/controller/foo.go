package controller

import (
	"github.com/labstack/echo/v4"
	"go-svc-tpl/app/response"
	"go-svc-tpl/model"
	"net/http"
)

// @tags     Foo
// @summary  A demo API
// @router   /foo [get]
// @produce  json
// @produce  json
// @param    foo  query     string  true  "Demo request parameter"
// @success  200  {object}  response.FooResponse
func Foo(c echo.Context) error {
	// just a demo
	var result model.Foo
	model.DB.Model(&model.Foo{}).First(&result)

	return c.JSON(http.StatusOK, response.FooResponse{
		Foo: c.QueryParam("foo"),
		Bar: 42,
	})

}
