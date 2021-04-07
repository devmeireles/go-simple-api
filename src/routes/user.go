package routes

import (
	router "go-backoffice-seller-api/src/http"
	"go-backoffice-seller-api/src/repositories/implementations"
	"go-backoffice-seller-api/src/useCase/readUser"

	"github.com/jinzhu/gorm"
)

func UserRoute(db *gorm.DB, httpRouter router.IRouter) router.IRouter {
	userRepository := implementations.NewUserRepository(db)
	readUserUseCase := readUser.NewReadUserUseCase(userRepository)
	readUserController := readUser.NewReadUserController(readUserUseCase)

	httpRouter.GET("/user/{id}", readUserController.GetUser)

	return httpRouter
}
