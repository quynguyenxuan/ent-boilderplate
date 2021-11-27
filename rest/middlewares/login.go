package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func RedirectToHomePageOnLogin(c *fiber.Ctx) error {
	// if auth.IsLoggedIn(c) {
	// 	return c.Redirect("/")
	// }
	return c.Next()
}
