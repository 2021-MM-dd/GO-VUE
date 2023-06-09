package main

import (
	"common/route"
	"fmt"
)

func main() {
	ruu := route.New()
	ruu.Use(Logger())
	ruu.GET("/v1/hello", func(ctx *route.Context) {
		ctx.String(200, "", "hello world")
	})
	ruu.GET("/v1/world", func(ctx *route.Context) {
		ctx.String(200, "", "hello world")
	})
	ruu.Run(":80")
}

func Logger() route.HandlerFunc {
	return func(c *route.Context) {
		fmt.Print("logger............")
		c.Next()
	}
}
