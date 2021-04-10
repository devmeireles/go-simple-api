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
	mailRepository implementations.IMailProvider
)

func NewCreateUserUseCase(repository implementations.IUserRepository, mail implementations.IMailProvider) ICreateUserUseCase {
	userRepository = repository
	mailRepository = mail
	return &createUserUseCase{}
}

func (createUserUseCase *createUserUseCase) Execute(user *entities.User) (*entities.User, error) {
	msgSubject := "Ownshop - Welcome " + user.Name
	mailRepository.SendMail(user.Email, user.Name, user.Language, msgSubject, "CREATE_ACCOUNT")

	return userRepository.CreateUser(user)
}
