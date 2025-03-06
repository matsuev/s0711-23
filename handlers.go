package main

import "github.com/gofiber/fiber/v2"

// RootHandler ...
func RootHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("hello")
}
