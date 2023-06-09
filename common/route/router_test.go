package route

import (
	"go-vue/common/constant"
	"testing"
)

func testget(t *testing.T) {
	ruu := New()
	ruu.GET("/test1", handle1)
	ruu.RUN("8888")
}

func handle1(ctx *Context) {
	ctx.String(int(constant.SUCCESS), "", "fucking jobless")
}
