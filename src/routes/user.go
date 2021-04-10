package routes

import (
	"go-backoffice-seller-api/src/repositories/implementations"
	"go-backoffice-seller-api/src/useCase/createUser"
	"go-backoffice-seller-api/src/useCase/readUser"
	router "go-backoffice-seller-api/src/utils"

	"github.com/jinzhu/gorm"
)

func UserRoute(db *gorm.DB, httpRouter router.IRouter) router.IRouter {
	userRepository := implementations.NewUserRepository(db)
	mailRepository := implementations.NewMailRepository()

	readUserUseCase := readUser.NewReadUserUseCase(userRepository)
	readUserController := readUser.NewReadUserController(readUserUseCase)

	createUserUseCase := createUser.NewCreateUserUseCase(userRepository, mailRepository)
	createUserController := createUser.NewCreateUserController(createUserUseCase)

	httpRouter.GET("/user/{id}", readUserController.Handler)
	httpRouter.POST("/user", createUserController.Handler)

	return httpRouter
}
