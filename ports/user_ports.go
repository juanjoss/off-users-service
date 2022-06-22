package ports

import (
	"net/http"

	"github.com/juanjoss/off-users-service/model"
)

type UserRepository interface {
	Register(*model.User, []*model.SSD) error
	AddProductToSSD(ssdId int, barcode string, quantity int) error
	RandomSSD() (*model.SSD, error)
}

type UserService interface {
	Register(*model.User, []*model.SSD) error
	AddProductToSSD(ssdId int, barcode string, quantity int) error
	RandomSSD() (*model.SSD, error)
}

type UserHandlers interface {
	Register(w http.ResponseWriter, r *http.Request)
	AddProductToSSD(w http.ResponseWriter, r *http.Request)
	RandomSSD(w http.ResponseWriter, r *http.Request)
}
