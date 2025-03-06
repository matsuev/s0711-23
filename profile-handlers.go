package main

import (
	"github.com/gofiber/fiber/v2"
)

// ProfileHandler ...
type ProfileHandler struct{}

// NewProfileHandler ...
func NewProfileHandler() CRUDHandler {
	return new(ProfileHandler)
}

// Read ...
func (u *ProfileHandler) Read() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Profile read")
	}
}

// Create ...
func (u *ProfileHandler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Profile create")
	}
}

// Update ...
func (u *ProfileHandler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Profile update")
	}
}

// Delete ...
func (u *ProfileHandler) Delete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Profile delete")
	}
}
