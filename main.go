package main

import (
	"os"

	"github.com/juanjoss/off-users-service/handlers"
	"github.com/juanjoss/off-users-service/repository"
	"github.com/juanjoss/off-users-service/server"
	"github.com/juanjoss/off-users-service/services"

	_ "github.com/juanjoss/off-users-service/docs"
)

func main() {
	// repositories
	ur := repository.NewUserRepository()

	// services
	us := services.NewUserService(ur)

	// handlers
	uh := handlers.NewUserHandlers(us)

	// server
	server.NewServer(uh).ListenAndServe(":" + os.Getenv("SERVICE_PORT"))
}
