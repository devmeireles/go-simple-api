package login

import (
	"encoding/json"
	"go-backoffice-seller-api/src/entities"
	"go-backoffice-seller-api/src/utils"
	"io/ioutil"
	"net/http"
)

type ILoginController interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type loginController struct{}

var (
	loginService ILoginUseCase
)

func NewLoginController(useCase ILoginUseCase) ILoginController {
	loginService = useCase
	return &loginController{}
}

func (loginController *loginController) Handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	login := entities.Login{}
	err = json.Unmarshal(body, &login)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	res, err := loginService.Execute(&login)

	if err != nil {
		utils.ResErr(w, err, http.StatusBadRequest)
		return
	}
	utils.ResSuc(w, res)

}
