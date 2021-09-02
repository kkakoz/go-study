package test

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"testing"
)

func TestHttp(t *testing.T) {

	app := fiber.New()

	app.Get("/echo", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello")
	})

	app.Get("/:name", func(ctx *fiber.Ctx) error {
		name := ctx.Params("name")
		return ctx.SendFile(fmt.Sprintf("./%s.xlsx", name))
	})

	app.Get("/temp/:name", func(ctx *fiber.Ctx) error {
		name := ctx.Params("name")
		return ctx.SendFile(fmt.Sprintf("./temp/%s.xlsx", name))
	})

	app.Get("/user/:id", func(ctx *fiber.Ctx) error {
		type ID struct {
			ID int `json:"id"`
		}
		id := &ID{}
		err := ctx.QueryParser(id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(ctx.ParamsInt("id"))
		return ctx.SendString(strconv.Itoa(id.ID))
	})

	app.Get("/stu/:id/name", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		return ctx.SendString(id)
	})

	app.Get("/stu/:ise_id", func(ctx *fiber.Ctx) error {
		return ctx.SendString("a")
	})

	app.Get("/temp", func(ctx *fiber.Ctx) error {
		return ctx.SendFile(fmt.Sprintf("./temp.xlsx"))
	})

	err := app.Listen(":10010")
	if err != nil {
		t.Fatal(err)
	}
}

