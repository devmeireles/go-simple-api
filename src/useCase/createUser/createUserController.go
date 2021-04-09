package createUser

import (
	"encoding/json"
	"go-backoffice-seller-api/src/entities"
	"go-backoffice-seller-api/src/utils"
	"io/ioutil"
	"net/http"
)

type ICreateUserController interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type createUserController struct{}

var (
	createUserService ICreateUserUseCase
)

func NewCreateUserController(useCase ICreateUserUseCase) ICreateUserController {
	createUserService = useCase
	return &createUserController{}
}

func (createUserController *createUserController) Handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	user := entities.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	validation := ValidateUser(&user)

	if validation != nil {
		utils.ResValidation(w, validation)
		return
	}

	user.Password = utils.HashAndSalt([]byte(user.Password))

	res, err := createUserService.Execute(&user)

	if err != nil {
		utils.ResErr(w, err, http.StatusBadRequest)
		return
	}
	utils.ResSuc(w, res)
}
