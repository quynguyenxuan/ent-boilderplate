package routes

import (
	"entgo.io/quynguyen-todo/rest/controllers"
	"entgo.io/quynguyen-todo/rest/middlewares"
	"github.com/gofiber/fiber/v2"
)

func WebAuthRoutes(App fiber.Router) {
	App.Get("/login", middlewares.RedirectToHomePageOnLogin, controllers.Login)
}
