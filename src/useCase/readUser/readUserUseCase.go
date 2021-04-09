package readUser

import (
	"go-backoffice-seller-api/src/entities"
	"go-backoffice-seller-api/src/repositories/implementations"
)

type IReadUserUseCase interface {
	Execute(id string) (*entities.User, error)
}

type readUserUseCase struct{}

var (
	userRepository implementations.IUserRepository
)

func NewReadUserUseCase(repository implementations.IUserRepository) IReadUserUseCase {
	userRepository = repository
	return &readUserUseCase{}
}

func (readUserUseCase *readUserUseCase) Execute(id string) (*entities.User, error) {
	return userRepository.GetUserById(id)
}
