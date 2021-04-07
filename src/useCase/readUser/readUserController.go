package readUser

import (
	"database/sql"
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
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	user, err := readUserService.GetUserService(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}
