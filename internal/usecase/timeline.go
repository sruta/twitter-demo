package usecase

import (
	"twitter-demo/internal/domain"
	"twitter-demo/internal/infraestructure/repository"
	"twitter-demo/pkg"
)

type ITimeline interface {
	GetTimeline(userID int64) ([]domain.Tweet, pkg.Error)
}

type Timeline struct {
	tweetRepository repository.ITweet
}

func NewTimeline(tweetRepository repository.ITweet) Timeline {
	return Timeline{
		tweetRepository: tweetRepository,
	}
}

func (t Timeline) GetTimeline(userID int64) ([]domain.Tweet, pkg.Error) {
	//Here we can add some business logic to process the tweets shown to the user
	//The famous "algorithm" that rules our lives should be here
	return t.tweetRepository.SelectForTimeline(userID)
}
