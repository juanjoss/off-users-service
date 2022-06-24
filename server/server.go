package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/juanjoss/off-users-service/ports"
)

const apiPrefix = "/api/users"

type Server struct {
	userHandlers ports.UserHandlers
	router       *mux.Router
	port         string
}

func NewServer(uh ports.UserHandlers) *Server {
	return &Server{
		userHandlers: uh,
		router:       mux.NewRouter().PathPrefix(apiPrefix).Subrouter(),
		port:         ":" + os.Getenv("SERVICE_PORT"),
	}
}

func (s *Server) RegisterRoutes() {
	s.router.HandleFunc("/register", s.userHandlers.Register).Methods(http.MethodPost)
	s.router.HandleFunc("/ssds/products", s.userHandlers.AddProductToSSD).Methods(http.MethodPost)
	s.router.HandleFunc("/ssds/random", s.userHandlers.RandomSSD).Methods(http.MethodGet)
}

func (s *Server) ListenAndServe() {
	s.RegisterRoutes()

	log.Println("Starting server at", s.port)
	log.Fatal(http.ListenAndServe(s.port, s.router))
}
