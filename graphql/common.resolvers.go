package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"entgo.io/quynguyen-todo/ent"
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

func (r *mutationResolver) CreateUser(ctx context.Context, user UserInput) (*ent.User, error) {
	client := ent.FromContext(ctx)
	return client.User.
		Create().
		SetStatus(user.Status).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetName(user.Name).
		SetEmail(user.Email).
		SetProviderID(user.ProviderID).
		SetProviderName(user.ProviderName).
		Save(ctx)
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
