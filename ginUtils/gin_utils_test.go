package ginUtils

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestServiceHandler(t *testing.T) {
	type A struct {
		X int
		Y int
	}
	type B string

	h := ServiceHandler(func(c *Context) []interface{} {
		return []Arg{A{123, 456}, B("xxx")}
	}, func(args ...interface{}) (result interface{}, err error) {
		a, ok := args[0].(A)
		if !ok {
			return nil, fmt.Errorf("Faield A")
		}
		b, ok := args[1].(B)
		if !ok {
			return nil, fmt.Errorf("Faield B")
		}
		fmt.Println(a, b)
		return []interface{}{a, b}, nil
	})
	h(&gin.Context{})
}
