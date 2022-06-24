package ports

import (
	"net/http"

	"github.com/juanjoss/off-users-service/model"
)

/*
	Interfaces
*/

type UserRepository interface {
	Register(RegisterRequest) error
	AddProductToSSD(AddProductToSsdRequest) error
	RandomSSD() (GetRandomSsdResponse, error)
}

type UserService interface {
	Register(RegisterRequest) error
	AddProductToSSD(AddProductToSsdRequest) error
	RandomSSD() (GetRandomSsdResponse, error)
}

type UserHandlers interface {
	Register(w http.ResponseWriter, r *http.Request)
	AddProductToSSD(w http.ResponseWriter, r *http.Request)
	RandomSSD(w http.ResponseWriter, r *http.Request)
}

/*
	Service models
*/

type RegisterRequest struct {
	User *model.User  `json:"user"`
	SSDs []*model.SSD `json:"devices"`
}

type AddProductToSsdRequest struct {
	SsdId    int    `json:"ssd_id"`
	Barcode  string `json:"barcode"`
	Quantity int    `json:"quantity"`
}

type GetRandomSsdResponse model.SSD
