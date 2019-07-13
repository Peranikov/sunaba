package graghql

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"net/http"
)

func Run() error {
	router := chi.NewRouter()

	router.Handle("/", handler.Playground("Dataloader", "/query"))
	router.Handle("/query", handler.GraphQL(
		NewExecutableSchema(Config{Resolvers: &Resolver{}}),
	))

	if err := http.ListenAndServe(":8082", router); err != nil {
		return err
	}

	return nil
}
