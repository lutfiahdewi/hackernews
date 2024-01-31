package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lutfiahdewi/hackernews/graph"
	"github.com/lutfiahdewi/hackernews/graph/generated"
	"github.com/lutfiahdewi/hackernews/internal/auth"
	// _ "github.com/glyphack/graphlq-golang/internal/auth"
	database "github.com/lutfiahdewi/hackernews/internal/pkg/db/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"	
	"github.com/go-chi/chi/v5"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	/*srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)*/

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	// defer database.CloseDB()
	database.Migrate()
	// server := handler.NewDefaultServer(hackernews.NewExecutableSchema(hackernews.Config{Resolvers: &hackernews.Resolver{}}))
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
