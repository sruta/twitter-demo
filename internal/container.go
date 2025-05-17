package internal

import (
	"fmt"
	"twitter-uala/internal/configs"
	"twitter-uala/internal/controller"
	"twitter-uala/internal/helpers"
	"twitter-uala/internal/repository/mysql"
	"twitter-uala/internal/service"
	"twitter-uala/pkg"
)

type Container struct {
	UserController controller.IUser
	AuthController controller.IAuth
}

func StartContainer() (*Container, error) {
	rdb, err := pkg.NewMySQL(configs.MySQLProd)
	if err != nil {
		fmt.Printf("Error creating rdb with %v", configs.MySQLProd)
		return nil, err
	}

	helpers.InitializeJWT(configs.JWTProd.Secret, configs.JWTProd.Expiration)

	userRepository := repository.NewUser(rdb)

	userService := service.NewUser(rdb, userRepository)
	userController := controller.NewUser(userService)

	authService := service.NewAuth(userRepository)
	authController := controller.NewAuth(authService)

	return &Container{
		UserController: userController,
		AuthController: authController,
	}, nil
}
