package services

import (
	"github.com/juanjoss/off-users-service/model"
	"github.com/juanjoss/off-users-service/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Register(user *model.User, ssds []*model.SSD) error {
	err := us.repo.Register(user, ssds)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) AddProductToSSD(ssdId int, barcode string, quantity int) error {
	err := us.repo.AddProductToSSD(ssdId, barcode, quantity)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) RandomSSD() (*model.SSD, error) {
	ssd, err := us.repo.RandomSSD()
	if err != nil {
		return ssd, err
	}

	return ssd, nil
}
