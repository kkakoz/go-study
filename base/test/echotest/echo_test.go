package echotest__test

import (
	"github.com/labstack/echo"
	"testing"
)

func TestEcho(t *testing.T) {
	e := echo.New()
	e.GET("/user/:user_id/name", func(ctx echo.Context) error {
		userId := ctx.Param("user_id")
		return ctx.String(200, userId + "name")
	})
	e.GET("/user/:id/str", func(ctx echo.Context) error {
		userId := ctx.Param("id")
		return ctx.String(200, userId + "str")
	})

	e.Start(":10235")

}
