package main

import (
	"log"

	"github.com/danielpomas/go-social/internal/env"
	"github.com/danielpomas/go-social/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	s := store.NewPostgresStorage(nil)

	app := &application{
		config: cfg,
		store:  s,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
