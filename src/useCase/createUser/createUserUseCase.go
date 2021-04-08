package createUser

import (
	"go-backoffice-seller-api/src/entities"
	"go-backoffice-seller-api/src/repositories/implementations"
)

type ICreateUserUseCase interface {
	Execute(user *entities.User) (*entities.User, error)
}

type createUserUseCase struct{}

var (
	userRepository implementations.IUserRepository
)

func NewCreateUserUseCase(repository implementations.IUserRepository) ICreateUserUseCase {
	userRepository = repository
	return &createUserUseCase{}
}

func (createUserUseCase *createUserUseCase) Execute(user *entities.User) (*entities.User, error) {
	return userRepository.CreateUser(user)
}
