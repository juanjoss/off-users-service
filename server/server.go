package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juanjoss/off-users-service/ports"
)

const apiPrefix = "/api/users"

type Server struct {
	userHandlers ports.UserHandlers
	router       *mux.Router
}

func NewServer(uh ports.UserHandlers) *Server {
	return &Server{
		userHandlers: uh,
		router:       mux.NewRouter().PathPrefix(apiPrefix).Subrouter(),
	}
}

func (s *Server) RegisterRoutes() {
	s.router.HandleFunc("/register", s.userHandlers.Register).Methods(http.MethodPost)
	s.router.HandleFunc("/ssds/products", s.userHandlers.AddProductToSSD).Methods(http.MethodPost)
	s.router.HandleFunc("/ssds/random", s.userHandlers.RandomSSD).Methods(http.MethodGet)
}

func (s *Server) ListenAndServe(addr string) {
	s.RegisterRoutes()

	log.Println("Starting server at", addr)
	if err := http.ListenAndServe(addr, s.router); err != nil {
		log.Fatal(err.Error())
	}
}
