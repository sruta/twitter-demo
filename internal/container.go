package internal

import (
	"fmt"
	"twitter-uala/internal/config"
	"twitter-uala/internal/infraestructure/repository"
	"twitter-uala/internal/interface/controller"
	"twitter-uala/internal/usecase"
	"twitter-uala/pkg"
)

type Container struct {
	UserController controller.IUser
	AuthController controller.IAuth
}

func StartContainer() (*Container, error) {
	rdb, err := pkg.NewMySQL(config.MySQLProd)
	if err != nil {
		fmt.Printf("Error creating rdb with %v", config.MySQLProd)
		return nil, err
	}

	pkg.InitializeJWT(config.JWTProd.Secret, config.JWTProd.Expiration)

	userRepository := repository.NewUser(rdb)
	userService := usecase.NewUser(rdb, userRepository)
	userController := controller.NewUser(userService)

	authService := usecase.NewAuth(userRepository)
	authController := controller.NewAuth(authService)

	return &Container{
		UserController: userController,
		AuthController: authController,
	}, nil
}
