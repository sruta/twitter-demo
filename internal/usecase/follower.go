package usecase

import (
	"twitter-uala/internal/domain"
	"twitter-uala/internal/infraestructure/repository"
	"twitter-uala/pkg"
)

type IFollower interface {
	Create(follower domain.Follower) (domain.Follower, pkg.Error)
}

type Follower struct {
	followerRepository repository.IFollower
	userRepository     repository.IUser
}

func NewFollower(followerRepository repository.IFollower, userRepository repository.IUser) Follower {
	return Follower{
		followerRepository: followerRepository,
		userRepository:     userRepository,
	}
}

func (f Follower) Create(follower domain.Follower) (domain.Follower, pkg.Error) {
	if follower.FollowerID == follower.FollowedID {
		return follower, pkg.NewGenericError("user cannot follow himself", nil)
	}

	_, err := f.userRepository.SelectByID(follower.FollowerID)
	if err != nil {
		return follower, err
	}

	_, err = f.userRepository.SelectByID(follower.FollowedID)
	if err != nil {
		return follower, err
	}

	repeatedFollower, err := f.followerRepository.SelectFollowerByIDs(follower.FollowerID, follower.FollowedID)
	if err != nil && !pkg.IsNotFound(err) {
		return follower, err
	}
	if repeatedFollower.FollowerID != 0 && repeatedFollower.FollowedID != 0 {
		return follower, pkg.NewGenericError("follower already exist", nil)
	}

	err = f.followerRepository.Insert(follower)
	if err != nil {
		return follower, err
	}

	return f.followerRepository.SelectFollowerByIDs(follower.FollowerID, follower.FollowedID)
}
