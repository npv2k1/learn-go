package main

import (
	"net/http"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}})))
	http.ListenAndServe(":8080", nil)
}
