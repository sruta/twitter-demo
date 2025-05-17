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
	FollowerController controller.IFollower
	UserController     controller.IUser
	AuthController     controller.IAuth
}

func StartContainer() (*Container, error) {
	rdb, err := pkg.NewMySQL(config.MySQLProd)
	if err != nil {
		fmt.Printf("Error creating rdb with %v", config.MySQLProd)
		return nil, err
	}

	pkg.InitializeJWT(config.JWTProd.Secret, config.JWTProd.Expiration)

	userRepository := repository.NewUser(rdb)
	userUsecase := usecase.NewUser(rdb, userRepository)
	userController := controller.NewUser(userUsecase)

	followerRepository := repository.NewFollower(rdb)
	followerUsecase := usecase.NewFollower(rdb, followerRepository, userRepository)
	followerController := controller.NewFollower(followerUsecase)

	authService := usecase.NewAuth(userRepository)
	authController := controller.NewAuth(authService)

	return &Container{
		FollowerController: followerController,
		UserController:     userController,
		AuthController:     authController,
	}, nil
}
