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
package main

import (
	"context"
	// "net/http"

	todo "entgo.io/quynguyen-todo/graphql"

	// "entgo.io/quynguyen-todo"
	"entgo.io/contrib/entgql"
	"entgo.io/quynguyen-todo/ent"
	"entgo.io/quynguyen-todo/ent/migrate"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"
	"go.uber.org/zap"

	_ "entgo.io/quynguyen-todo/ent/runtime"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var cli struct {
		Addr  string `name:"address" default:":8081" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}
	kong.Parse(&cli)

	log, _ := zap.NewDevelopment()
	// client, err := ent.Open(
	// 	"postgres",
	// 	"host=127.0.0.1 port=5432 user=postgres dbname=ent password=123456 sslmode=disable",
	// )
	client, err := ent.Open(
		"sqlite3",
		// "file:ent?mode=memory&cache=shared&_fk=1",
		"file:sqlite.db?cache=shared&_fk=1",
	)
	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("running schema migration", zap.Error(err))
	}

	srv := handler.NewDefaultServer(todo.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	if cli.Debug {
		srv.Use(&debug.Tracer{})
	}
	app := fiber.New()

	app.All("/query", func(c *fiber.Ctx) error {
		// var r http.Request
		ctx := c.Context()

		fasthttpH := fasthttpadaptor.NewFastHTTPHandlerFunc(srv.ServeHTTP)
		fasthttpH(ctx)
		return nil
	})

	app.All("/", func(c *fiber.Ctx) error {
		ctx := c.Context()
		fasthttpH := fasthttpadaptor.NewFastHTTPHandlerFunc(playground.Handler("Todo", "/query"))
		fasthttpH(ctx)
		return nil
	})

	app.Listen(cli.Addr)

	// http.Handle("/",
	// 	playground.Handler("Todo", "/query"),
	// )
	// http.Handle("/query", srv)

	// if err := http.ListenAndServe(cli.Addr, nil); err != nil {
	// 	log.Error("http server terminated", zap.Error(err))
	// }

	log.Info("listening on", zap.String("address", cli.Addr))
}