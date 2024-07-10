package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type List struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var lists []List = []List{}

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello, World!")
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Get("/lists", func(c *fiber.Ctx) error {
		fmt.Println(lists)
		return c.JSON(fiber.Map{
			"lists": lists,
		})
	})

	app.Post("/lists", func(c *fiber.Ctx) error {
		list := new(List)
		if err := c.BodyParser(list); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		lists = append(lists, *list)
		return c.JSON(fiber.Map{
			"message": "List added successfully",
			"lists":   lists,
		})
	})

	app.Listen(":3000")
}
