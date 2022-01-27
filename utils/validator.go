package utils

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

type CustomValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (c *CustomValidator) Validate(i interface{}) error {
	c.lazyInit()
	return c.validate.Struct(i)
}

func (c *CustomValidator) lazyInit() {
	c.once.Do(func() {
		c.validate = validator.New()
	})
}
