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
	"time"

	"entgo.io/quynguyen-todo/ent"
	"entgo.io/quynguyen-todo/ent/category"
	"entgo.io/quynguyen-todo/ent/product"
	"entgo.io/quynguyen-todo/ent/schema/schematype"
	"entgo.io/quynguyen-todo/ent/todo"
)

type (
	Category656363463View struct {
		ID     int                        `json:"id,omitempty"`
		Text   string                     `json:"text,omitempty"`
		Status category.Status            `json:"status,omitempty"`
		Config *schematype.CategoryConfig `json:"config,omitempty"`
	}
	Category656363463Views []*Category656363463View
)

func NewCategory656363463View(e *ent.Category) *Category656363463View {
	return &Category656363463View{
		ID:     e.ID,
		Text:   e.Text,
		Status: e.Status,
		Config: e.Config,
	}
}

func NewCategory656363463Views(es []*ent.Category) Category656363463Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Category656363463Views, len(es))
	for i, e := range es {
		r[i] = NewCategory656363463View(e)
	}
	return r
}

type (
	Product1899176864View struct {
		ID        int            `json:"id,omitempty"`
		CreatedAt time.Time      `json:"created_at,omitempty"`
		Status    product.Status `json:"status,omitempty"`
		Priority  int            `json:"priority,omitempty"`
		Text      string         `json:"text,omitempty"`
	}
	Product1899176864Views []*Product1899176864View
)

func NewProduct1899176864View(e *ent.Product) *Product1899176864View {
	return &Product1899176864View{
		ID:        e.ID,
		CreatedAt: e.CreatedAt,
		Status:    e.Status,
		Priority:  e.Priority,
		Text:      e.Text,
	}
}

func NewProduct1899176864Views(es []*ent.Product) Product1899176864Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Product1899176864Views, len(es))
	for i, e := range es {
		r[i] = NewProduct1899176864View(e)
	}
	return r
}

type (
	Todo2548332322View struct {
		ID        int         `json:"id,omitempty"`
		CreatedAt time.Time   `json:"created_at,omitempty"`
		Status    todo.Status `json:"status,omitempty"`
		Priority  int         `json:"priority,omitempty"`
		Text      string      `json:"text,omitempty"`
		Blob      []byte      `json:"blob,omitempty"`
	}
	Todo2548332322Views []*Todo2548332322View
)

func NewTodo2548332322View(e *ent.Todo) *Todo2548332322View {
	return &Todo2548332322View{
		ID:        e.ID,
		CreatedAt: e.CreatedAt,
		Status:    e.Status,
		Priority:  e.Priority,
		Text:      e.Text,
		Blob:      e.Blob,
	}
}

func NewTodo2548332322Views(es []*ent.Todo) Todo2548332322Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Todo2548332322Views, len(es))
	for i, e := range es {
		r[i] = NewTodo2548332322View(e)
	}
	return r
}

type (
	VerySecret1653553545View struct {
		ID       int    `json:"id,omitempty"`
		Password string `json:"password,omitempty"`
	}
	VerySecret1653553545Views []*VerySecret1653553545View
)

func NewVerySecret1653553545View(e *ent.VerySecret) *VerySecret1653553545View {
	return &VerySecret1653553545View{
		ID:       e.ID,
		Password: e.Password,
	}
}

func NewVerySecret1653553545Views(es []*ent.VerySecret) VerySecret1653553545Views {
	if len(es) == 0 {
		return nil
	}
	r := make(VerySecret1653553545Views, len(es))
	for i, e := range es {
		r[i] = NewVerySecret1653553545View(e)
	}
	return r
}
