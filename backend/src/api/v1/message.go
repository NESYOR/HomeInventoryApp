package v1

import "github.com/gofiber/fiber"

func SayHello(c *fiber.Ctx) error{
	return c.SendString("Hello")
}