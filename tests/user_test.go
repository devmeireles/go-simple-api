package tests

import (
	"bytes"
	"encoding/json"
	"go-backoffice-seller-api/src/entities"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
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

type CurrentUser struct {
	userID string
}

var primaryUser = entities.User{
	Name:     faker.Name(),
	Email:    faker.Email(),
	Password: faker.Password(),
}

var wrongUser = entities.User{
	Name:     faker.Name(),
	Password: faker.Password(),
}

var currentUser CurrentUser

func TestUserModule(t *testing.T) {
	t.Run("POST", func(t *testing.T) {
		t.Run("It should create a user", func(t *testing.T) {
			var user = primaryUser

			userSave, _ := json.Marshal(user)

			req, _ := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(userSave))

			response := ExecuteRequest(req)
			parsedBody := ParseBody(response)

			userID := parsedBody.Data.(map[string]interface{})["id"]
			currentUser.userID = userID.(string)

			assert.True(t, parsedBody.Success)
			assert.Equal(t, http.StatusOK, response.Code)
		})

		t.Run("It shouldn't create a user because there are missing some fields", func(t *testing.T) {
			var user = wrongUser

			userSave, _ := json.Marshal(user)

			req, _ := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(userSave))

			response := ExecuteRequest(req)
			parsedBody := ParseBody(response)

			assert.False(t, parsedBody.Success)
			assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
		})

		t.Run("It shouldn't create a user because the user already exists", func(t *testing.T) {
			var user = primaryUser

			userSave, _ := json.Marshal(user)

			req, _ := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(userSave))

			response := ExecuteRequest(req)
			parsedBody := ParseBody(response)

			assert.False(t, parsedBody.Success)
			assert.Equal(t, http.StatusBadRequest, response.Code)
		})
	})

	t.Run("GET", func(t *testing.T) {
		t.Run("It should return a user", func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/api/v1/user/"+currentUser.userID, nil)
			response := ExecuteRequest(req)
			parsedBody := ParseBody(response)

			assert.True(t, parsedBody.Success)
			assert.Equal(t, http.StatusOK, response.Code)
			assert.Empty(t, parsedBody.Message)
			assert.Contains(t, parsedBody.Data, "name")
		})

		t.Run("It shouldn't return a user because he doesn't exist", func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/api/v1/user/51", nil)
			response := ExecuteRequest(req)
			parsedBody := ParseBody(response)

			assert.False(t, parsedBody.Success)
			assert.Equal(t, http.StatusNotFound, response.Code)
			assert.NotEmpty(t, parsedBody.Message)
			assert.Empty(t, parsedBody.Data)
		})
	})
}
