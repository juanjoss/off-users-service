package services

import (
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

func (us *UserService) Register(request ports.RegisterRequest) error {
	err := us.repo.Register(request)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) AddProductToSSD(request ports.AddProductToSsdRequest) error {
	err := us.repo.AddProductToSSD(request)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) RandomSSD() (ports.GetRandomSsdResponse, error) {
	response, err := us.repo.RandomSSD()
	if err != nil {
		return response, err
	}

	return response, nil
}
