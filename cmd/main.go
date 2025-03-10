package main

import (
	"fmt"
	"http/test/configs"
	"http/test/internal/auth"
	"http/test/internal/link"
	"http/test/internal/stat"
	"http/test/internal/user"
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
	userRepository := user.NewUserRepository(dbInstance)
	statRepository := stat.NewStatRepository(dbInstance)

	// Services
	authService := auth.NewAuthService(userRepository)

	// Handler
	auth.NewHelloHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		StatRepository: statRepository,
		Config:         conf,
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
