package repository

import (
	"database/sql"
	"errors"
	"twitter-uala/internal/domain"
	"twitter-uala/pkg"
)

type ITweet interface {
	SelectByID(id int64) (domain.Tweet, pkg.Error)
	Insert(tweet domain.Tweet) (domain.Tweet, pkg.Error)
	Update(tweet domain.Tweet) (domain.Tweet, pkg.Error)
}

type Tweet struct {
	rdb *pkg.MySQL
}

func NewTweet(rdb *pkg.MySQL) Tweet {
	return Tweet{
		rdb: rdb,
	}
}

func (t Tweet) SelectByID(id int64) (domain.Tweet, pkg.Error) {
	var result domain.Tweet

	row := t.rdb.QueryRow("select id, user_id, text, created_at, updated_at from tweet where id = ?", id)
	err := row.Scan(&result.ID, &result.UserID, &result.Text, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, pkg.NewDBNotFoundError("tweet", err)
		}
		return result, pkg.NewDBScanFatalError("tweet", err)
	}

	return result, nil
}

func (t Tweet) Insert(tweet domain.Tweet) (domain.Tweet, pkg.Error) {
	query := "insert into tweet(user_id, text, created_at, updated_at) values (?,?, now(), now())"
	result, err := t.rdb.Exec(query, tweet.UserID, tweet.Text)
	if err != nil {
		return tweet, pkg.NewDBFatalError("insert tweet into", err)
	}

	tweet.ID, err = result.LastInsertId()
	if err != nil {
		return tweet, pkg.NewDBFatalError("insert tweet into", err)
	}

	return tweet, nil
}

func (t Tweet) Update(tweet domain.Tweet) (domain.Tweet, pkg.Error) {
	var err error
	query := "update tweet set text = ?, updated_at = now() where id = ?"
	_, err = t.rdb.Exec(query, tweet.Text, tweet.ID)
	if err != nil {
		return tweet, pkg.NewDBFatalError("update tweet in", err)
	}

	return tweet, nil
}
