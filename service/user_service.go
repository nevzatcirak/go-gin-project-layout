package service

import (
	"github.com/nevzatcirak/go-gin-poc/domain"
)

type UserService interface {
	Save(domain.User) error
	Update(domain.User) error
	Delete(domain.User) error
	Find(id uint64) domain.User
	FindAll() []domain.User
}

type userService struct {
	repository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) UserService {
	return &userService{
		repository: userRepository,
	}
}

func (service *userService) Save(user domain.User) error {
	service.repository.Save(user)
	return nil
}

func (service *userService) Update(user domain.User) error {
	service.repository.Update(user)
	return nil
}

func (service *userService) Delete(user domain.User) error {
	service.repository.Delete(user)
	return nil
}

func (service *userService) Find(id uint64) domain.User {
	return service.repository.Find(id)
}

func (service *userService) FindAll() []domain.User {
	return service.repository.FindAll()
}
