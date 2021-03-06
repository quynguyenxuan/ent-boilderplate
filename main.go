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
	todo "entgo.io/quynguyen-todo/graphql"

	// "entgo.io/quynguyen-todo"
	"entgo.io/contrib/entgql"
	"entgo.io/quynguyen-todo/ent"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"

	_ "entgo.io/quynguyen-todo/ent/runtime"
	controllers "entgo.io/quynguyen-todo/rest/controllers"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	_ "entgo.io/quynguyen-todo/docs"
	// elk "entgo.io/quynguyen-todo/ent/http"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/adaptor/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

type HTTPError struct {
	fiber.Error
}

// @title Fiber API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email quynguyen@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// @BasePath /api
func main() {
	var cli struct {
		Addr  string `name:"address" default:":8081" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
		Prod  bool   `name:"prod" default:"false" help:"Enable prefork in Production"`
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
	// postgresql://localhost:5432/postgres?user=postgres&password=postgrespassword&sslmode=disable
	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
	}
	//auto migrate
	// if err := client.Schema.Create(
	// 	context.Background(),
	// 	migrate.WithGlobalUniqueID(true),
	// ); err != nil {
	// 	log.Fatal("running schema migration", zap.Error(err))
	// }

	srv := handler.NewDefaultServer(todo.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	if cli.Debug {
		srv.Use(&debug.Tracer{})
	}

	privateKey := controllers.GenerateKey()

	app := fiber.New(fiber.Config{
		Prefork: cli.Prod, // go run app.go -prod
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.All("/graphql", adaptor.HTTPHandlerFunc(srv.ServeHTTP))
	app.All("/playground", adaptor.HTTPHandlerFunc(playground.Handler("Todo", "/graphql")))

	apiGroup := app.Group("/api")
	userGroup := apiGroup.Group("/user")
	userGroup.Use(jwtware.New(jwtware.Config{
		// SigningKey: []byte("secret"),
		SigningMethod: "RS256",
		SigningKey:    privateKey.Public(),
	}))
	// apiHandle := controllers.NewHandler(client, log)
	apiGroup.Post("/login", controllers.Login)
	userGroup.Get("/", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		// name := claims["name"].(string)
		return c.JSON(claims)
		// return c.SendString("Welcome " + name)
	})
	// api.NewHandler(client, log).Mount(apiGroup)
	// productsRouter := api.Group("/product")
	// elk.NewProductHandler(client, log).Mount(productsRouter, elk.ProductRoutes)

	// todosRouter := api.Group("/todo")
	// elk.NewProductHandler(client, log).Mount(todosRouter, elk.TodoRoutes)

	app.Get("/swagger/*", swagger.Handler) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http:/localhost:8081/doc.json",
		DeepLinking: true,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		// OAuth: &swagger.OAuthConfig{
		// 	AppName:  "OAuth Provider",
		// 	ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		// },
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8081/swagger/oauth2-redirect.html",
	}))
	// app.Static("/static", "./static/public")

	app.Listen(cli.Addr)

	log.Info("listening on", zap.String("address", cli.Addr))
}
