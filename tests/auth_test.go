package tests

import (
	"bytes"
	"encoding/json"
	"go-backoffice-seller-api/src/entities"
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

type CreatedUser struct {
	userID   string
	email    string
	password string
}

type LoginData struct {
	Email    string
	Password string
}

var createUser = entities.User{
	Name:     faker.Name(),
	Email:    faker.Email(),
	Password: faker.Password(),
}

var createdUser CreatedUser

func TestAuthModule(t *testing.T) {
	t.Run("POST", func(t *testing.T) {
		t.Run("It should create a user to be used on login", func(t *testing.T) {
			var user = createUser

			userSave, _ := json.Marshal(user)

			req, _ := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(userSave))

			response := ExecuteRequest(req)
			parsedBody := ParseBody(response)

			userID := parsedBody.Data.(map[string]interface{})["id"]
			createdUser.userID = userID.(string)
			createdUser.email = user.Email
			createdUser.password = user.Password

			assert.True(t, parsedBody.Success)
			assert.Equal(t, http.StatusOK, response.Code)
		})

		t.Run("It should login and return a token", func(t *testing.T) {
			var user = LoginData{
				Email:    createdUser.email,
				Password: createdUser.password,
			}

			login, _ := json.Marshal(user)

			req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(login))

			response := ExecuteRequest(req)
			parsedBody := ParseBody(response)

			assert.True(t, parsedBody.Success)
			assert.Equal(t, http.StatusOK, response.Code)
		})

		t.Run("It should login because the user doesn't exist", func(t *testing.T) {
			var user = LoginData{
				Email:    faker.Email(),
				Password: createdUser.password,
			}

			login, _ := json.Marshal(user)

			req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(login))

			response := ExecuteRequest(req)
			parsedBody := ParseBody(response)

			assert.False(t, parsedBody.Success)
			assert.Equal(t, http.StatusBadRequest, response.Code)
		})

		t.Run("It should login because the password is incorrect", func(t *testing.T) {
			var user = LoginData{
				Email:    createdUser.email,
				Password: faker.Password(),
			}

			login, _ := json.Marshal(user)

			req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(login))

			response := ExecuteRequest(req)
			parsedBody := ParseBody(response)

			assert.False(t, parsedBody.Success)
			assert.Equal(t, http.StatusBadRequest, response.Code)
		})

		t.Run("It shouldn't create a user because there are missing some fields", func(t *testing.T) {
			var user = LoginData{}

			login, _ := json.Marshal(user)

			req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(login))

			response := ExecuteRequest(req)
			parsedBody := ParseBody(response)

			assert.False(t, parsedBody.Success)
			assert.Equal(t, http.StatusBadRequest, response.Code)
		})
	})
}
