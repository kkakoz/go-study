package test

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestJson(t *testing.T) {
	a := []int{}
	data := []byte("[1,2,3]")
	err := json.Unmarshal(data, &a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)

	curT, err := time.Parse("2006-01-02", "2001-01-01")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(curT)
	fmt.Println(curT.Unix())
}

type IA interface {
	AA()
}

type A struct {

}

func (receiver A) AA() {

}

func TestT(t *testing.T) {
	ias := make([]IA, 0)
	ias = append(ias, A{})
	handleA(ias)
}

func handleA(as []IA) {
	for _, v := range as {
		v.AA()
	}
}

type Context interface {
	context.Context
}


var str = "goroutine 917 [running]:\\nruntime/debug.Stack(0x1074700, 0xc0006efc80, 0x11cd4f5)\\n\\t/usr/local/go/src/runtime/debug/stack.go:24 +0x9f\\ngit.kelexuexi.com/ixue-rpc/go-sdk/framework.Exception.func1.1(0xc0001b4400)\\n\\t/go/pkg/mod/git.kelexuexi.com/ixue-rpc/go-sdk@v0.7.8/framework/middleware.go:73 +0x3d7\\npanic(0x106b020, 0xc0005f6888)\\n\\t/usr/local/go/src/runtime/panic.go:965 +0x1b9\\nmain/controller.(*Context).ErrHandle(...)\\n\\t/www/src/controller/base.go:121\\nmain/controller/edu.DispatchController.TeacherDetail(0xc0001b4400)\\n\\t/www/src/controller/edu/dispatch.go:56 +0x4b2\\ngithub.com/gin-gonic/gin.(*Context).Next(...)\\n\\t/go/pkg/mod/github.com/gin-gonic/gin@v1.7.1/context.go:165\\nmain/bootstrap.useMiddleware.func1(0xc0001b4400)\\n\\t/www/src/bootstrap/middleware.go:14 +0x62\\ngithub.com/gin-gonic/gin.(*Context).Next(...)\\n\\t/go/pkg/mod/github.com/gin-gonic/gin@v1.7.1/context.go:165\\ngit.kelexuexi.com/ixue-rpc/go-sdk/framework.Exception.func1(0xc0001b4400)\\n\\t/go/pkg/mod/git.kelexuexi.com/ixue-rpc/go-sdk@v0.7.8/framework/middleware.go:90 +0x63\\ngithub.com/gin-gonic/gin.(*Context).Next(...)\\n\\t/go/pkg/mod/github.com/gin-gonic/gin@v1.7.1/context.go:165\\ngit.kelexuexi.com/ixue-rpc/go-sdk/framework.Base.func1(0xc0001b4400)\\n\\t/go/pkg/mod/git.kelexuexi.com/ixue-rpc/go-sdk@v0.7.8/framework/middleware.go:29 +0x266\\ngithub.com/gin-gonic/gin.(*Context).Next(...)\\n\\t/go/pkg/mod/github.com/gin-gonic/gin@v1.7.1/context.go:165\\ngithub.com/gin-gonic/gin.(*Engine).handleHTTPRequest(0xc0003364e0, 0xc0001b4400)\\n\\t/go/pkg/mod/github.com/gin-gonic/gin@v1.7.1/gin.go:489 +0x2aa\\ngithub.com/gin-gonic/gin.(*Engine).ServeHTTP(0xc0003364e0, 0x135e408, 0xc0006395e0, 0xc0005c1100)\\n\\t/go/pkg/mod/github.com/gin-gonic/gin@v1.7.1/gin.go:445 +0x15c\\nnet/http.serverHandler.ServeHTTP(0xc0003d8000, 0x135e408, 0xc0006395e0, 0xc0005c1100)\\n\\t/usr/local/go/src/net/http/server.go:2887 +0xa3\\nnet/http.(*conn).serve(0xc000221360, 0x1361220, 0xc000360d40)\\n\\t/usr/local/go/src/net/http/server.go:1952 +0x8cd\\ncreated by net/http.(*Server).Serve\\n\\t/usr/local/go/src/net/http/server.go:3013 +0x39b\\n"

func TestA(t *testing.T) {
	split := strings.Split(str, "\\n")
	for _, v := range split {

		fmt.Println(strings.Trim(v, "\\t"))
	}
}


type customFunc func(is ...int)

func a(c customFunc)  {
	c()
}


type TTT struct {
	Name string
}

func TestTTT(t *testing.T)  {

}
