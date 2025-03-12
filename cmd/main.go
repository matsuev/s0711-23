package main

import (
	"log"

	"github.com.matsuev/s0711-23/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	api := fiber.New()

	svc := service.New()

	api.Get("/", svc.RootHandler)

	api.Get("/api/user", svc.UserRead)

	api.Get("/api/post", svc.PostRead)

	if err := api.Listen(":7080"); err != nil {
		log.Fatalln(err)
	}
}
