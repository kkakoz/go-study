package echotest__test

import (
	"github.com/labstack/echo"
	"testing"
)

func TestEcho(t *testing.T) {
	e := echo.New()
	e.GET("/user/:user_id/name", func(ctx echo.Context) error {
		userId := ctx.Param("user_id")
		return ctx.String(200, userId+"name")
	})
	e.GET("/user/:id/str", func(ctx echo.Context) error {
		type ID struct {
			Id string
		}
		id := &ID{}
		err := (&echo.DefaultBinder{}).Bind(id, ctx)
		if err != nil {
			return err
		}
		return ctx.String(200, "s"+id.Id)
	})

	e.Start(":10235")

}
