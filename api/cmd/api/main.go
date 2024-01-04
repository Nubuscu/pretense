package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"nubuscu/pretense"
	"nubuscu/pretense/ent"
	"nubuscu/pretense/ent/migrate"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	// n.b. ssl hardcoded off for local dev.
	connStr := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"))
	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on :8081.
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	}).Handler)

	srv := handler.NewDefaultServer(pretense.NewSchema(client))
	router.Handle("/",
		playground.Handler("Api", "/query"),
	)
	router.Handle("/query", srv)

	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatal("http server terminated", err)
	}
}
