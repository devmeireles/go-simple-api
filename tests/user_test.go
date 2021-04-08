package tests

import (
	"bytes"
	"encoding/json"
	"go-backoffice-seller-api/src/entities"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	httpRouter.SERVEHTTP(rr, req)

	return rr
}

func ParseBody(content *httptest.ResponseRecorder) entities.ResponseMsg {
	parsedRes := entities.ResponseMsg{}
	body, _ := ioutil.ReadAll(content.Body)
	json.Unmarshal(body, &parsedRes)

	return parsedRes

}

func TestUserModule(t *testing.T) {

	t.Run("It should create a user", func(t *testing.T) {
		var user = entities.User{
			Name:     "Gabriel Meireles",
			Email:    "dev.meireles@gmail.com",
			Password: "aVeryStrongPassword",
		}

		userSave, _ := json.Marshal(user)

		req, _ := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(userSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	// t.Run("It should return a user", func(t *testing.T) {
	// 	req, _ := http.NewRequest("GET", "/api/v1/user/1", nil)
	// 	response := ExecuteRequest(req)
	// 	parsedBody := ParseBody(response)

	// 	assert.True(t, parsedBody.Success)
	// 	assert.Equal(t, http.StatusOK, response.Code)
	// 	assert.Empty(t, parsedBody.Message)
	// 	assert.Contains(t, parsedBody.Data, "name")
	// })

	// t.Run("It shouldn't return a user because he doesn't exist", func(t *testing.T) {
	// 	req, _ := http.NewRequest("GET", "/api/v1/user/51", nil)
	// 	response := ExecuteRequest(req)
	// 	parsedBody := ParseBody(response)

	// 	assert.False(t, parsedBody.Success)
	// 	assert.Equal(t, http.StatusNotFound, response.Code)
	// 	assert.NotEmpty(t, parsedBody.Message)
	// 	assert.Empty(t, parsedBody.Data)
	// })
}
