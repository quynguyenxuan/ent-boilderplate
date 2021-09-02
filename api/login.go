package api

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

var (
	// Obviously, this is just a test example. Do not do this in production.
	// In production, you would have the private key and public key pair generated
	// in advance. NEVER add a private key to any GitHub repo.
	privateKey *rsa.PrivateKey
)

func GenerateKey() *rsa.PrivateKey {
	rng := rand.Reader
	var err error
	privateKey, err = rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}
	return privateKey
}

func (h *Handler) Mount(r fiber.Router) {
	r.Post("/login", h.Login)
}

type LoginRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

//LoginApi
func (h *Handler) Login(c *fiber.Ctx) error {
	l := h.log.With(zap.String("method", "Login"))
	// Get the post data.
	// r := HttpRequest(c)

	d := new(LoginRequest)
	if err := c.BodyParser(d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		return c.Status(400).SendString("invalid json string")
	}
	if *d.Username != "xuanquy" || *d.Password != "123456" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create token
	// token := jwt.New(jwt.SigningMethodES256)
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(privateKey)
	// t, err := token.SignedString([]byte("secret"))
	if err != nil {
		// fmt.Println(err)
		l.Error("error create token", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})

	// Save the data.
	// b := h.client.Product.Create()
	// if d.Username != nil {
	// 	b.SetText(*d.Username)
	// }
	// if d.Username != nil {
	// 	b.SetText(*d.Username)
	// }
	// e, err := b.Save(r.Context())
	// if err != nil {
	// 	switch {
	// 	default:
	// 		l.Error("could not create toy", zap.Error(err))
	// 		c.Status(fiber.StatusInternalServerError).SendString("Serve Error")
	// 	}
	// 	return nil
	// }
	// Reload entry.
	// q := h.client.Product.Query().Where(e.ID(e.ID))
	// e, err = q.Only(r.Context())
	// if err != nil {
	// 	switch {
	// 	case ent.IsNotFound(err):
	// 		msg := StripEntError(err)
	// 		l.Info(msg, zap.Error(err), zap.Int("id", e.ID))
	// 		c.Status(404).SendString(msg)
	// 	case ent.IsNotSingular(err):
	// 		msg := StripEntError(err)
	// 		l.Error(msg, zap.Error(err), zap.Int("id", e.ID))
	// 		c.Status(400).SendString(msg)
	// 	default:
	// 		l.Error("could not read toy", zap.Error(err), zap.Int("id", e.ID))
	// 		c.Status(fiber.StatusInternalServerError).SendString("Serve Error")
	// 	}
	// 	return nil
	// }
	// l.Info("toy rendered", zap.Int("id", e.ID))
	// return c.JSON(e)
}
