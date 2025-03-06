package main

import (
	"log"
	"sync/atomic"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var UserMap = make(map[int]UserModel)
var UserID int64 = 0

// UserHandler ...
type UserHandler struct {
	validate *validator.Validate
}

// NewUserHandler ...
func NewUserHandler() CRUDHandler {
	return &UserHandler{
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

// Read ...
func (u *UserHandler) Read() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		intID := ctx.QueryInt("id")

		if intID < 1 {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		user, ok := UserMap[intID]
		if !ok {
			return ctx.SendStatus(fiber.StatusNotFound)
		}

		return ctx.JSON(user)
	}
}

// Create ...
func (u *UserHandler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var user UserModel
		if err := ctx.BodyParser(&user); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		if err := u.validate.Struct(user); err != nil {
			log.Println(err)
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		user.ID = int(atomic.AddInt64(&UserID, 1))

		UserMap[user.ID] = user

		return ctx.JSON(user)
	}
}

// Update ...
func (u *UserHandler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("user update")
	}
}

// Delete ...
func (u *UserHandler) Delete() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		intID := ctx.QueryInt("id")

		if intID < 1 {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		delete(UserMap, intID)

		return ctx.SendString("user delete")
	}
}
