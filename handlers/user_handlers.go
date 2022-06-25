package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/juanjoss/off-users-service/ports"
)

type UserHandlers struct {
	userService ports.UserService
}

func NewUserHandlers(userService ports.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (uh *UserHandlers) Register(w http.ResponseWriter, r *http.Request) {
	var request ports.RegisterRequest

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = uh.userService.Register(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(http.StatusOK)
}

func (uh *UserHandlers) AddProductToSSD(w http.ResponseWriter, r *http.Request) {
	var request ports.AddProductToSsdRequest

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = uh.userService.AddProductToSSD(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(http.StatusOK)
}

func (uh *UserHandlers) RandomSSD(w http.ResponseWriter, r *http.Request) {
	response, err := uh.userService.RandomSSD()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
