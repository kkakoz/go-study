package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello")
	})

	err := app.Listen(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
