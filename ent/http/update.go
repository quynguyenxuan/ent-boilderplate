// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"

	"entgo.io/quynguyen-todo/ent"
	"entgo.io/quynguyen-todo/ent/category"
	"entgo.io/quynguyen-todo/ent/product"
	"entgo.io/quynguyen-todo/ent/todo"
	"entgo.io/quynguyen-todo/ent/verysecret"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"
)

// Update updates a given ent.Category and saves the changes to the database.
func (h CategoryHandler) Update(c *fiber.Ctx) error {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", c.Params("id")), zap.Error(err))
		return c.Status(400).SendString("id must be an integer greater zero")
	}
	// Get the post data.
	d := new(CategoryUpdateRequest)
	var r http.Request
	fasthttpadaptor.ConvertRequest(c.Context(), &r, true)

	if err := c.BodyParser(d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		return c.Status(400).SendString("invalid json string")
	}
	// Validate the data.
	errs := make(map[string]string)
	if d.Text == nil {
		errs["text"] = `missing required field: "text"`
	} else if err := category.TextValidator(*d.Text); err != nil {
		errs["text"] = err.Error()
	}
	if d.Status == nil {
		errs["status"] = `missing required field: "status"`
	} else if err := category.StatusValidator(*d.Status); err != nil {
		errs["status"] = err.Error()
	}
	if len(errs) > 0 {
		l.Info("validation failed", zapFields(errs)...)
		return c.Status(400).JSON(errs)

	}
	// Save the data.
	b := h.client.Category.UpdateOneID(id)
	if d.Text != nil {
		b.SetText(*d.Text)
	}
	if d.Status != nil {
		b.SetStatus(*d.Status)
	}
	if d.Config != nil {
		b.SetConfig(*d.Config)
	}
	if d.Todos != nil {
		b.ClearTodos().AddTodoIDs(d.Todos...)
	}
	// Store in database.
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			c.Status(404).SendString(msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			c.Status(400).SendString(msg)
		default:
			l.Error("could-not-update-category", zap.Error(err), zap.Int("id", id))
			c.Status(fiber.StatusInternalServerError).SendString("Serve Error")
		}
		return nil
	}
	// Reload entry.
	q := h.client.Category.Query().Where(category.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			c.Status(404).SendString(msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			c.Status(400).SendString(msg)
		default:
			l.Error("could-not-read-category", zap.Error(err), zap.Int("id", id))
			c.Status(fiber.StatusInternalServerError).SendString("Serve Error")
		}
		return nil
	}
	l.Info("category rendered", zap.Int("id", e.ID))
	return c.JSON(NewCategory656363463View(e))
}

// Update updates a given ent.Product and saves the changes to the database.
func (h ProductHandler) Update(c *fiber.Ctx) error {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", c.Params("id")), zap.Error(err))
		return c.Status(400).SendString("id must be an integer greater zero")
	}
	// Get the post data.
	d := new(ProductUpdateRequest)
	var r http.Request
	fasthttpadaptor.ConvertRequest(c.Context(), &r, true)

	if err := c.BodyParser(d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		return c.Status(400).SendString("invalid json string")
	}
	// Validate the data.
	errs := make(map[string]string)
	// if d.CreatedAt == nil {
	// 	errs["created_at"] = `missing required field: "created_at"`
	// }
	if d.Status == nil {
		errs["status"] = `missing required field: "status"`
	} else if err := product.StatusValidator(*d.Status); err != nil {
		errs["status"] = err.Error()
	}
	if d.Priority == nil {
		errs["priority"] = `missing required field: "priority"`
	}
	if d.Text == nil {
		errs["text"] = `missing required field: "text"`
	} else if err := product.TextValidator(*d.Text); err != nil {
		errs["text"] = err.Error()
	}
	if len(errs) > 0 {
		l.Info("validation failed", zapFields(errs)...)
		return c.Status(400).JSON(errs)

	}
	// Save the data.
	b := h.client.Product.UpdateOneID(id)
	if d.Status != nil {
		b.SetStatus(*d.Status)
	}
	if d.Priority != nil {
		b.SetPriority(*d.Priority)
	}
	if d.Text != nil {
		b.SetText(*d.Text)
	}
	// Store in database.
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			c.Status(404).SendString(msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			c.Status(400).SendString(msg)
		default:
			l.Error("could-not-update-product", zap.Error(err), zap.Int("id", id))
			c.Status(fiber.StatusInternalServerError).SendString("Serve Error")
		}
		return nil
	}
	// Reload entry.
	q := h.client.Product.Query().Where(product.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			c.Status(404).SendString(msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			c.Status(400).SendString(msg)
		default:
			l.Error("could-not-read-product", zap.Error(err), zap.Int("id", id))
			c.Status(fiber.StatusInternalServerError).SendString("Serve Error")
		}
		return nil
	}
	l.Info("product rendered", zap.Int("id", e.ID))
	return c.JSON(NewProduct1899176864View(e))
}

// Update updates a given ent.Todo and saves the changes to the database.
func (h TodoHandler) Update(c *fiber.Ctx) error {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", c.Params("id")), zap.Error(err))
		return c.Status(400).SendString("id must be an integer greater zero")
	}
	// Get the post data.
	d := new(TodoUpdateRequest)
	var r http.Request
	fasthttpadaptor.ConvertRequest(c.Context(), &r, true)

	if err := c.BodyParser(d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		return c.Status(400).SendString("invalid json string")
	}
	// Validate the data.
	errs := make(map[string]string)
	// if d.CreatedAt == nil {
	// 	errs["created_at"] = `missing required field: "created_at"`
	// }
	if d.Status == nil {
		errs["status"] = `missing required field: "status"`
	} else if err := todo.StatusValidator(*d.Status); err != nil {
		errs["status"] = err.Error()
	}
	if d.Priority == nil {
		errs["priority"] = `missing required field: "priority"`
	}
	if d.Text == nil {
		errs["text"] = `missing required field: "text"`
	} else if err := todo.TextValidator(*d.Text); err != nil {
		errs["text"] = err.Error()
	}
	if len(errs) > 0 {
		l.Info("validation failed", zapFields(errs)...)
		return c.Status(400).JSON(errs)

	}
	// Save the data.
	b := h.client.Todo.UpdateOneID(id)
	if d.Status != nil {
		b.SetStatus(*d.Status)
	}
	if d.Priority != nil {
		b.SetPriority(*d.Priority)
	}
	if d.Text != nil {
		b.SetText(*d.Text)
	}
	if d.Blob != nil {
		b.SetBlob(*d.Blob)
	}
	if d.Parent != nil {
		b.SetParentID(*d.Parent)

	}
	if d.Children != nil {
		b.ClearChildren().AddChildIDs(d.Children...)
	}
	if d.Category != nil {
		b.SetCategoryID(*d.Category)

	}
	if d.Secret != nil {
		b.SetSecretID(*d.Secret)

	}
	// Store in database.
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			c.Status(404).SendString(msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			c.Status(400).SendString(msg)
		default:
			l.Error("could-not-update-todo", zap.Error(err), zap.Int("id", id))
			c.Status(fiber.StatusInternalServerError).SendString("Serve Error")
		}
		return nil
	}
	// Reload entry.
	q := h.client.Todo.Query().Where(todo.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			c.Status(404).SendString(msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			c.Status(400).SendString(msg)
		default:
			l.Error("could-not-read-todo", zap.Error(err), zap.Int("id", id))
			c.Status(fiber.StatusInternalServerError).SendString("Serve Error")
		}
		return nil
	}
	l.Info("todo rendered", zap.Int("id", e.ID))
	return c.JSON(NewTodo2548332322View(e))
}

// Update updates a given ent.VerySecret and saves the changes to the database.
func (h VerySecretHandler) Update(c *fiber.Ctx) error {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", c.Params("id")), zap.Error(err))
		return c.Status(400).SendString("id must be an integer greater zero")
	}
	// Get the post data.
	d := new(VerySecretUpdateRequest)
	var r http.Request
	fasthttpadaptor.ConvertRequest(c.Context(), &r, true)

	if err := c.BodyParser(d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		return c.Status(400).SendString("invalid json string")
	}
	// Save the data.
	b := h.client.VerySecret.UpdateOneID(id)
	if d.Password != nil {
		b.SetPassword(*d.Password)
	}
	// Store in database.
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			c.Status(404).SendString(msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			c.Status(400).SendString(msg)
		default:
			l.Error("could-not-update-very-secret", zap.Error(err), zap.Int("id", id))
			c.Status(fiber.StatusInternalServerError).SendString("Serve Error")
		}
		return nil
	}
	// Reload entry.
	q := h.client.VerySecret.Query().Where(verysecret.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			c.Status(404).SendString(msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			c.Status(400).SendString(msg)
		default:
			l.Error("could-not-read-very-secret", zap.Error(err), zap.Int("id", id))
			c.Status(fiber.StatusInternalServerError).SendString("Serve Error")
		}
		return nil
	}
	l.Info("very-secret rendered", zap.Int("id", e.ID))
	return c.JSON(NewVerySecret1653553545View(e))
}
