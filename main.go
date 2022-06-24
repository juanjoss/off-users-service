package main

import (
	"github.com/juanjoss/off-users-service/handlers"
	"github.com/juanjoss/off-users-service/repository"
	"github.com/juanjoss/off-users-service/server"
	"github.com/juanjoss/off-users-service/services"

	_ "github.com/juanjoss/off-users-service/docs"
)

func main() {
	ur := repository.NewUserRepository()
	us := services.NewUserService(ur)
	uh := handlers.NewUserHandlers(us)
	server.NewServer(uh).ListenAndServe()
}
