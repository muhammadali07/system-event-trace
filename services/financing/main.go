package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/graph-gophers/graphql-go"
)

// Schema GraphQL
const schemaString = `
	type Query {
		hello: String!
	}

	schema {
		query: Query
	}
`

// Resolver untuk menangani query
type Resolver struct{}

// Metode resolver untuk menangani query hello
func (r *Resolver) Hello() string {
	return "Hello, world!"
}

func main() {
	// Setup Fiber
	app := fiber.New()

	// Setup GraphQL Schema
	schema := graphql.MustParseSchema(schemaString, &Resolver{})

	// Handler untuk endpoint GraphQL
	app.All("/graphql", func(c *fiber.Ctx) error {
		ctx := context.Background()
		result := schema.Exec(ctx, "", "", c.Request().Body())
		return c.Status(http.StatusOK).SendString(result.Encode())
	})

	// Handler untuk Playground GraphQL
	app.Get("/", func(c *fiber.Ctx) error {
		playground := playgroundHandler("/graphql", "/subscriptions")
		return c.Status(http.StatusOK).SendString(playground)
	})

	// Mulai server
	server := &http.Server{
		Addr:         ":3000",
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Server GraphQL berjalan di http://localhost:3000")
	log.Fatal(server.ListenAndServe())
}

func playgroundHandler(endpoint, subscriptionEndpoint string) string {
	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
		<head>
			<link rel="stylesheet" href="https://unpkg.com/graphql-playground-react/build/static/css/index.css" />
			<script src="https://unpkg.com/react/umd/react.production.min.js"></script>
			<script src="https://unpkg.com/react-dom/umd/react-dom.production.min.js"></script>
			<script src="https://unpkg.com/graphql-playground-react/build/static/js/middleware.js"></script>
		</head>
		<body>
			<div id="root"></div>
			<script type="text/javascript">
				const root = document.getElementById('root')
				const wsProto = location.protocol == 'https:' ? 'wss:' : 'ws:'
				const subscriptionEndpoint = "%s"
				const subscriptionUrl = wsProto + '//' + location.host + subscriptionEndpoint
				const fetcher = GraphiQL.createFetcher('%s')

				ReactDOM.render(
					React.createElement(GraphQLPlayground, { 
						endpoint: '%s', 
						subscriptionEndpoint: subscriptionUrl,
						fetcher: fetcher 
					}),
					root
				)
			</script>
		</body>
	</html>
	`, subscriptionEndpoint, endpoint, endpoint)
}
