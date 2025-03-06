package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// CRUDHandler ...
type CRUDHandler interface {
	Create() fiber.Handler
	Read() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
}

func main() {
	app := fiber.New()

	initCRUDHandler(app, "/api/user", NewUserHandler())
	// initCRUDHandler(app, "/api/profile", NewProfileHandler(pgconn))

	if err := app.Listen(":7080"); err != nil {
		log.Fatalln(err)
	}
}

// initCRUDHandler ...
func initCRUDHandler(app *fiber.App, path string, h CRUDHandler) {
	app.Get(path, h.Read())
	app.Post(path, h.Create())
	app.Put(path, h.Update())
	app.Delete(path, h.Delete())
}
