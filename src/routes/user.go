package routes

import (
	"go-backoffice-seller-api/src/repositories/implementations"
	"go-backoffice-seller-api/src/useCase/readUser"
	router "go-backoffice-seller-api/src/utils"

	"github.com/jinzhu/gorm"
)

func UserRoute(db *gorm.DB, httpRouter router.IRouter) router.IRouter {
	userRepository := implementations.NewUserRepository(db)
	readUserUseCase := readUser.NewReadUserUseCase(userRepository)
	readUserController := readUser.NewReadUserController(readUserUseCase)

	httpRouter.GET("/user/{id}", readUserController.Handler)

	return httpRouter
}
