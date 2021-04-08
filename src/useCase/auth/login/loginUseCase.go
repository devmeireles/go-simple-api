package login

import (
	"errors"
	"go-backoffice-seller-api/src/entities"
	"go-backoffice-seller-api/src/repositories/implementations"
	"go-backoffice-seller-api/src/utils"
)

type ILoginUseCase interface {
	Execute(user *entities.Login) (*entities.LoggedUser, error)
}

type loginUseCase struct{}

var (
	userRepository implementations.IUserRepository
)

func NewLoginUseCase(repository implementations.IUserRepository) ILoginUseCase {
	userRepository = repository
	return &loginUseCase{}
}

func (loginUseCase *loginUseCase) Execute(login *entities.Login) (*entities.LoggedUser, error) {
	currentUser, err := userRepository.GetUserByEmail(login.Email)

	if err != nil {
		return nil, err
	}

	if !utils.ComparePasswords(currentUser.Password, []byte(login.Password)) {
		return nil, errors.New("incorrect password")
	}

	token, _ := utils.CreateToken(currentUser.Email)

	loggedUser := entities.LoggedUser{}

	loggedUser.Email = currentUser.Email
	loggedUser.Name = currentUser.Name
	loggedUser.AccountType = currentUser.AccountType
	loggedUser.Language = currentUser.Language
	loggedUser.Token = token

	return &loggedUser, err
}
