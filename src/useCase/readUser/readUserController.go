package readUser

import (
	"database/sql"
	"go-backoffice-seller-api/src/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type IReadUserController interface {
	GetUser(w http.ResponseWriter, r *http.Request)
}

type userController struct{}

var (
	readUserService IReadUserUseCase
)

func NewReadUserController(service IReadUserUseCase) IReadUserController {
	readUserService = service
	return &userController{}
}

func (userController *userController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ResErr(w, err, http.StatusBadRequest)
		return
	}
	user, err := readUserService.GetUserService(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.ResErr(w, err, http.StatusNotFound)
		default:
			utils.ResErr(w, err, http.StatusInternalServerError)
		}
		return
	}
	utils.ResSuc(w, user)
}
