//go:build wireinject
// +build wireinject

package main

import (
	"github.com/SamedArslan28/gopost/internal/config"
	"github.com/SamedArslan28/gopost/internal/database"
	"github.com/SamedArslan28/gopost/internal/handler"
	"github.com/SamedArslan28/gopost/internal/repository"
	"github.com/SamedArslan28/gopost/internal/service"
	"github.com/google/wire"
)

// InitializeApp tells Wire how to build the server, accepting the app config as input.
func InitializeApp(cfg config.Config) (*Server, error) {
	wire.Build(
		// Config provider
		provideDatabaseDsn,

		// Database provider
		database.ConnectDB,

		// Repository providers
		repository.NewUserRepository,
		repository.NewPostRepository,

		// Service providers
		service.NewUserService,
		service.NewPostService,

		// Handler providers
		provideUserHandler,
		providePostHandler,

		// Server provider
		NewServer,
	)

	return nil, nil
}

func provideDatabaseDsn(cfg config.Config) string {
	return cfg.DatabaseURL
}

func provideUserHandler(userService *service.UserService) *handler.UserHandler {
	return handler.NewUserHandler(userService)
}

func providePostHandler(postService *service.PostService) *handler.PostHandler {
	return handler.NewPostHandler(postService)
}
