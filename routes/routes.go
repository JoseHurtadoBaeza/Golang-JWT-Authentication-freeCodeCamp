package routes

import (
	"github.com/JoseHurtadoBaeza/React-Golang-JWT-Authentication-freeCodeCamp/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

}
