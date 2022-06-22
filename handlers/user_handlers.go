package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/juanjoss/off-users-service/model"
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

type UserRegistrationRequest struct {
	User    *model.User  `json:"user"`
	Devices []*model.SSD `json:"devices"`
}

func (uh *UserHandlers) Register(w http.ResponseWriter, r *http.Request) {
	var request UserRegistrationRequest

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = uh.userService.Register(request.User, request.Devices)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(http.StatusOK)
}

type AddProductToUserSSDRequest struct {
	SsdId    int    `json:"ssd_id"`
	Barcode  string `json:"barcode"`
	Quantity int    `json:"n_products"`
}

func (uh *UserHandlers) AddProductToSSD(w http.ResponseWriter, r *http.Request) {
	var request AddProductToUserSSDRequest

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = uh.userService.AddProductToSSD(request.SsdId, request.Barcode, request.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(http.StatusOK)
}

func (uh *UserHandlers) RandomSSD(w http.ResponseWriter, r *http.Request) {
	ssd, err := uh.userService.RandomSSD()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ssd)
}
