//go:build wireinject
// +build wireinject

// This file is purely an instruction manual for the `wire` command.
// It tells Wire how to build the application by listing the necessary providers.

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
// It references the helper providers located in your providers.go file.
func InitializeApp(cfg config.Config) (*Server, error) {
	wire.Build(
		// Helper providers (from providers.go)
		provideDatabaseDsn,
		provideUserHandler,

		// Core component providers (from the internal package)
		database.ConnectDB,
		repository.NewUserRepository,
		service.NewUserService,

		// The final server provider
		NewServer,
	)

	// This return is a placeholder that Wire will replace.
	return nil, nil
}

func provideDatabaseDsn(cfg config.Config) string {
	return cfg.DatabaseURL
}

// provideUserHandler is a helper provider.
// Your original code was: handler.NewUserHandler(*userService)
// This function replicates that logic so Wire can use it.
func provideUserHandler(userService *service.UserService) handler.UserHandler {
	userHandlerPointer := handler.NewUserHandler(*userService)
	return *userHandlerPointer
}
