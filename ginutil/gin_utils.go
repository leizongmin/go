package ginutil

import (
	"github.com/gin-gonic/gin"
)

type H = gin.H
type Context = gin.Context
type Engine = gin.Engine

func Default() *gin.Engine {
	return gin.Default()
}

func New() *gin.Engine {
	return gin.New()
}

// 自定义响应处理函数
var ResponseHandler func(c *Context, data interface{}, err error)

func init() {
	ResponseHandler = func(c *Context, data interface{}, err error) {
		if err != nil {
			c.JSON(500, H{"ok": false, "error": err.Error()})
		} else {
			c.JSON(200, H{"ok": true, "data": data})
		}
	}
}

// 响应失败结果
func ResponseError(c *Context, err error) {
	ResponseHandler(c, nil, err)
}

// 响应成功结果
func ResponseOk(c *Context, data interface{}) {
	ResponseHandler(c, data, nil)
}

// 抽象服务处理接口
func ServiceHandler(parseParams func(c *Context) []interface{}, serviceHandler func(args ...interface{}) (result interface{}, err error)) func(c *Context) {
	return func(c *Context) {
		result, err := serviceHandler(parseParams(c)...)
		if err != nil {
			ResponseError(c, err)
		} else {
			ResponseOk(c, result)
		}
	}
}

type Arg = interface{}
