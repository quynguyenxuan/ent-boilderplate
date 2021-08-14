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
	"entgo.io/quynguyen-todo/ent/todo"
	"github.com/go-chi/chi/v5"
	"github.com/mailru/easyjson"
	"github.com/masseelch/render"
	"go.uber.org/zap"
)

// Todos fetches the ent.todos attached to the ent.Category
// identified by a given url-parameter from the database and renders it to the client.
func (h CategoryHandler) Todos(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Todos"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		render.BadRequest(w, r, "id must be an integer greater zero")
		return
	}
	// Create the query to fetch the todos attached to this category
	q := h.client.Category.Query().Where(category.ID(id)).QueryTodos()
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			render.BadRequest(w, r, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			render.BadRequest(w, r, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching todos from db", zap.Error(err))
		render.InternalServerError(w, r, nil)
		return
	}
	l.Info("todos rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewTodo2548332322Views(es), w)
}

// Parent fetches the ent.parent attached to the ent.Todo
// identified by a given url-parameter from the database and renders it to the client.
func (h TodoHandler) Parent(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Parent"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		render.BadRequest(w, r, "id must be an integer greater zero")
		return
	}
	// Create the query to fetch the parent attached to this todo
	q := h.client.Todo.Query().Where(todo.ID(id)).QueryParent()
	e, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			render.NotFound(w, r, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			render.BadRequest(w, r, msg)
		default:
			l.Error("could-not-read-todo", zap.Error(err), zap.Int("id", id))
			render.InternalServerError(w, r, nil)
		}
		return
	}
	l.Info("todo rendered", zap.Int("id", e.ID))
	easyjson.MarshalToHTTPResponseWriter(NewTodo2548332322View(e), w)
}

// Children fetches the ent.children attached to the ent.Todo
// identified by a given url-parameter from the database and renders it to the client.
func (h TodoHandler) Children(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Children"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		render.BadRequest(w, r, "id must be an integer greater zero")
		return
	}
	// Create the query to fetch the children attached to this todo
	q := h.client.Todo.Query().Where(todo.ID(id)).QueryChildren()
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			render.BadRequest(w, r, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			render.BadRequest(w, r, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching todos from db", zap.Error(err))
		render.InternalServerError(w, r, nil)
		return
	}
	l.Info("todos rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewTodo2548332322Views(es), w)
}

// Category fetches the ent.category attached to the ent.Todo
// identified by a given url-parameter from the database and renders it to the client.
func (h TodoHandler) Category(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Category"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		render.BadRequest(w, r, "id must be an integer greater zero")
		return
	}
	// Create the query to fetch the category attached to this todo
	q := h.client.Todo.Query().Where(todo.ID(id)).QueryCategory()
	e, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			render.NotFound(w, r, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			render.BadRequest(w, r, msg)
		default:
			l.Error("could-not-read-todo", zap.Error(err), zap.Int("id", id))
			render.InternalServerError(w, r, nil)
		}
		return
	}
	l.Info("category rendered", zap.Int("id", e.ID))
	easyjson.MarshalToHTTPResponseWriter(NewCategory656363463View(e), w)
}

// Secret fetches the ent.secret attached to the ent.Todo
// identified by a given url-parameter from the database and renders it to the client.
func (h TodoHandler) Secret(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Secret"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		render.BadRequest(w, r, "id must be an integer greater zero")
		return
	}
	// Create the query to fetch the secret attached to this todo
	q := h.client.Todo.Query().Where(todo.ID(id)).QuerySecret()
	e, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			render.NotFound(w, r, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			render.BadRequest(w, r, msg)
		default:
			l.Error("could-not-read-todo", zap.Error(err), zap.Int("id", id))
			render.InternalServerError(w, r, nil)
		}
		return
	}
	l.Info("very-secret rendered", zap.Int("id", e.ID))
	easyjson.MarshalToHTTPResponseWriter(NewVerySecret1653553545View(e), w)
}