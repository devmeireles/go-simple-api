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

type readUserController struct{}

var (
	readUserService IReadUserUseCase
)

func NewReadUserController(useCase IReadUserUseCase) IReadUserController {
	readUserService = useCase
	return &readUserController{}
}

func (readUserController *readUserController) Handler(w http.ResponseWriter, r *http.Request) {
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
