package readUser

import (
	"go-backoffice-seller-api/src/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type IReadUserController interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type userController struct{}

var (
	readUserService IReadUserUseCase
)

func NewReadUserController(useCase IReadUserUseCase) IReadUserController {
	readUserService = useCase
	return &userController{}
}

func (userController *userController) Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ResErr(w, err, http.StatusBadRequest)
		return
	}
	user, err := readUserService.Execute(id)
	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}
	utils.ResSuc(w, user)
}
