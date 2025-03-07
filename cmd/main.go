package main

import (
	"fmt"
	"http/test/configs"
	"http/test/internal/auth"
	"http/test/internal/link"
	"http/test/pkg/db"
	"http/test/pkg/middleware"
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

	// Middleware
	stack := middleware.Chain(middleware.CORS, middleware.Logging)

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: stack(router),
	}

	fmt.Println("Server starting on port 8081")
	server.ListenAndServe()
}
