package main

import (
	"fmt"
	"http/test/configs"
	"http/test/internal/auth"
	"http/test/internal/link"
	"http/test/internal/stat"
	"http/test/internal/user"
	"http/test/pkg/db"
	"http/test/pkg/event"
	"http/test/pkg/middleware"
	"net/http"
)

func App() http.Handler {
	conf := configs.LoadConfig()
	dbInstance := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := link.NewLinkRepository(dbInstance)
	userRepository := user.NewUserRepository(dbInstance)
	statRepository := stat.NewStatRepository(dbInstance)

	// Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus:       eventBus,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})

	go statService.AddClick()

	// Middlewares
	stack := middleware.Chain(middleware.CORS, middleware.Logging)

	return stack(router)
}

func main() {
	app := App()

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: app,
	}

	fmt.Println("Server starting on port 8081")
	server.ListenAndServe()
}
