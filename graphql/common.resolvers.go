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

package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	// "fmt"
	"time"

	"entgo.io/quynguyen-todo/ent"
	"entgo.io/quynguyen-todo/ent/todo"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, todo TodoInput) (*ent.Todo, error) {
	client := ent.FromContext(ctx)
	return client.Todo.
		Create().
		SetStatus(todo.Status).
		SetNillablePriority(todo.Priority).
		SetText(todo.Text).
		SetNillableParentID(todo.Parent).
		Save(ctx)
	// panic(fmt.Err?orf("not implemented"))
}

func (r *mutationResolver) ClearTodos(ctx context.Context) (int, error) {
	client := ent.FromContext(ctx)
	return client.Todo.
		Delete().
		Exec(ctx)
	// panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProduct(ctx context.Context, product ProductInput) (*ent.Product, error) {
	client := ent.FromContext(ctx)
	client.Todo.Create().SetStatus(todo.StatusCompleted).SetText(product.Text).SetNillablePriority(product.Priority).SetCreatedAt(time.Now()).Save(ctx)
	return client.Product.
		Create().
		SetStatus(product.Status).
		SetNillablePriority(product.Priority).
		SetText(product.Text).
		SetCreatedAt(time.Now()).
		Save(ctx)
	// panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ClearProducts(ctx context.Context) (int, error) {
	client := ent.FromContext(ctx)
	return client.Product.Delete().Exec(ctx)
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)

	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TodoOrder, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
			ent.WithTodoFilter(where.Filter),
		)
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.ProductOrder, where *ent.ProductWhereInput) (*ent.ProductConnection, error) {
	return r.client.Product.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithProductOrder(orderBy),
			ent.WithProductFilter(where.Filter),
		)
	// panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
