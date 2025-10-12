//go:build wireinject
// +build wireinject

// The above build tags tell the normal `go build` command to ignore this file.
// It's only used as input for the `wire` command-line tool.

package main

import (
	"github.com/google/wire"

	"github.com/SamedArslan28/gopost/internal/database"
	"github.com/SamedArslan28/gopost/internal/handler"
	"github.com/SamedArslan28/gopost/internal/repository"
	"github.com/SamedArslan28/gopost/internal/service"
)

func InitializeApp(dbDsn string) (*Server, error) {
	wire.Build(
		database.ConnectDB,

		repository.NewUserRepository,
		service.NewUserService,
		provideUserHandler,

		NewServer,
	)
	return nil, nil
}

// provideUserHandler is a helper provider.
// Your original code was: handler.NewUserHandler(*userService)
// This function replicates that logic so Wire can use it.
func provideUserHandler(userService *service.UserService) handler.UserHandler {
	// 1. Create the handler. `NewUserHandler` returns a pointer (*UserHandler).
	userHandlerPointer := handler.NewUserHandler(*userService)

	// 2. Dereference the pointer to get the actual value (UserHandler).
	// This matches the function's required return type.
	return *userHandlerPointer
}
