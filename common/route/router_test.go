package route

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	ruu := New()
	ruu.Use(Logger())
	ruu.GET("/v1/hello", func(ctx *Context) {
		ctx.String(200, "hello %s\n", "john")
	})
	ruu.GET("/v1/world", func(ctx *Context) {
		ctx.String(200, "hello %s\n", "alex")
	})
	ruu.Run(":80")
}

func Logger() HandlerFunc {
	return func(c *Context) {
		fmt.Println("logger............")
		c.Next()
		fmt.Println("hahahaha")
	}
}
