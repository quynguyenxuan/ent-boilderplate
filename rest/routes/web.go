package routes

import (
	"github.com/gofiber/fiber/v2"
)

func WebRoutes(web fiber.Router) {
	WebAuthRoutes(web)
}
