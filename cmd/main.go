package main

import (
	"fmt"
	"http/test/configs"
	"http/test/internal/auth"
	"http/test/internal/link"
	"http/test/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	dbInstance := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(dbInstance)

	// Handler
	auth.NewHelloHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: router,
	}

	fmt.Println("Server starting on port 8081")
	server.ListenAndServe()
}
