package internal

import (
	"fmt"
	"twitter-uala/internal/config"
	"twitter-uala/internal/infraestructure/repository"
	"twitter-uala/internal/interfaces/controller"
	"twitter-uala/internal/usecase"
	"twitter-uala/pkg"
)

type Container struct {
	TimelineController controller.ITimeline
	TweetController    controller.ITweet
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

	tweetRepository := repository.NewTweet(rdb)
	tweetUsecase := usecase.NewTweet(tweetRepository)
	tweetController := controller.NewTweet(tweetUsecase)

	userRepository := repository.NewUser(rdb)
	userUsecase := usecase.NewUser(userRepository)
	userController := controller.NewUser(userUsecase)

	followerRepository := repository.NewFollower(rdb)
	followerUsecase := usecase.NewFollower(followerRepository, userRepository)
	followerController := controller.NewFollower(followerUsecase)

	timelineUsecase := usecase.NewTimeline(tweetRepository)
	timelineController := controller.NewTimeline(timelineUsecase)

	authService := usecase.NewAuth(userRepository)
	authController := controller.NewAuth(authService)

	return &Container{
		TweetController:    tweetController,
		FollowerController: followerController,
		UserController:     userController,
		TimelineController: timelineController,
		AuthController:     authController,
	}, nil
}
