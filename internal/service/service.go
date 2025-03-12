package service

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Service ...
type Service struct{}

// New ...
func New() *Service {
	return new(Service)
}

// RootHandler ...
func (s *Service) RootHandler(ctx *fiber.Ctx) error {
	log.Println("request")
	return ctx.SendString("Hello!!!")
}

// UserRead ..
func (s *Service) UserRead(ctx *fiber.Ctx) error {
	log.Println("user")
	return ctx.SendString("User data")
}

// PostRead ...
func (s *Service) PostRead(ctx *fiber.Ctx) error {
	log.Println("post")
	return ctx.SendString("Post data")
}
