package createUser

import (
	"go-backoffice-seller-api/src/entities"

	"github.com/thedevsaddam/govalidator"
)

func ValidateUser(user *entities.User) map[string]interface{} {
	rules := govalidator.MapData{
		"name":     []string{"required"},
		"email":    []string{"required"},
		"password": []string{"required"},
	}

	opts := govalidator.Options{
		Data:  user,
		Rules: rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		return map[string]interface{}{"error": e}
	}

	return nil
}
