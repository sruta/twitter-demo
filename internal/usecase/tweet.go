package usecase

import (
	"twitter-uala/internal/domain"
	"twitter-uala/internal/infraestructure/repository"
	"twitter-uala/pkg"
)

type ITweet interface {
	Create(tweet domain.Tweet) (domain.Tweet, pkg.Error)
	Update(tweet domain.Tweet) (domain.Tweet, pkg.Error)
	SearchByID(id int64) (domain.Tweet, pkg.Error)
}

type Tweet struct {
	rdb             pkg.DB
	tweetRepository repository.ITweet
}

func NewTweet(rdb pkg.DB, tweetRepository repository.ITweet) Tweet {
	return Tweet{
		rdb:             rdb,
		tweetRepository: tweetRepository,
	}
}

func (t Tweet) Create(tweet domain.Tweet) (domain.Tweet, pkg.Error) {
	tweet, err := t.tweetRepository.Insert(tweet)
	if err != nil {
		return tweet, err
	}

	return t.tweetRepository.SelectByID(tweet.ID)
}

func (t Tweet) Update(tweet domain.Tweet) (domain.Tweet, pkg.Error) {
	dbTweet, err := t.tweetRepository.SelectByID(tweet.ID)
	if err != nil {
		return tweet, err
	}

	if dbTweet.UserID != tweet.UserID {
		return tweet, pkg.NewForbiddenError("user not authorized", nil)
	}

	dbTweet.Text = tweet.Text

	dbTweet, err = t.tweetRepository.Update(dbTweet)
	if err != nil {
		return tweet, err
	}

	return t.tweetRepository.SelectByID(dbTweet.ID)
}

func (t Tweet) SearchByID(id int64) (domain.Tweet, pkg.Error) {
	return t.tweetRepository.SelectByID(id)
}
