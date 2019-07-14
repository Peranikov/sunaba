package graphql

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

const (
	addr = ":8082"
)

func Run() error {
	router := chi.NewRouter()

	router.Handle("/", handler.Playground("Dataloader", "/query"))
	router.Handle("/query", handler.GraphQL(
		NewExecutableSchema(Config{Resolvers: &Resolver{}}),
	))

	log.Printf("started graghql to listen http://localhost%v\n", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		return err
	}

	return nil
}
