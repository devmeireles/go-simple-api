package routes

import (
	"go-backoffice-seller-api/src/repositories/implementations"
	"go-backoffice-seller-api/src/useCase/auth/login"
	router "go-backoffice-seller-api/src/utils"

	"github.com/jinzhu/gorm"
)

func AuthRoute(db *gorm.DB, httpRouter router.IRouter) router.IRouter {
	userRepository := implementations.NewUserRepository(db)

	loginUseCase := login.NewLoginUseCase(userRepository)
	loginController := login.NewLoginController(loginUseCase)

	httpRouter.POST("/login", loginController.Handler)

	return httpRouter
}
