package controllers

import (
	"context"
	"net/http"
	"strings"

	"entgo.io/quynguyen-todo/ent"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"
)

type handler struct{}

// handler has some convenience methods used on node-handlers.
type Handler struct {
	handler

	client *ent.Client
	log    *zap.Logger
}

func NewHandler(c *ent.Client, l *zap.Logger) *Handler {
	return &Handler{
		client: c,
		log:    l.With(zap.String("handler", "Handler")),
	}
}

// Bitmask to configure which routes to register.
type Routes uint32

func (rs Routes) Has(r Routes) bool { return rs&r != 0 }

func StripEntError(err error) string {
	return strings.TrimPrefix(err.Error(), "ent: ")
}

func ZapFields(errs map[string]string) []zap.Field {
	if errs == nil || len(errs) == 0 {
		return nil
	}
	r := make([]zap.Field, 0)
	for k, v := range errs {
		r = append(r, zap.String(k, v))
	}
	return r
}

func HttpRequest(c *fiber.Ctx) http.Request {
	var r http.Request
	fasthttpadaptor.ConvertRequest(c.Context(), &r, true)
	return r
}

func HttpContext(c *fiber.Ctx) context.Context {
	var r http.Request
	fasthttpadaptor.ConvertRequest(c.Context(), &r, true)
	return r.Context()
}
