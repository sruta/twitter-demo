package repository

import (
	"database/sql"
	"errors"
	"twitter-uala/internal/domain"
	"twitter-uala/pkg"
)

type IFollower interface {
	Insert(follower domain.Follower) pkg.Error
	SelectFollowerByIDs(followerID int64, followedID int64) (domain.Follower, pkg.Error)
}

type Follower struct {
	rdb *pkg.MySQL
}

func NewFollower(rdb *pkg.MySQL) IFollower {
	return &Follower{
		rdb: rdb,
	}
}

func (f *Follower) Insert(follower domain.Follower) pkg.Error {
	query := "insert into follower(follower_id, followed_id, created_at) values (?,?, now())"
	_, err := f.rdb.Exec(query, follower.FollowerID, follower.FollowedID)
	if err != nil {
		return pkg.NewDBFatalError("insert follower into", err)
	}

	return nil
}

func (f *Follower) SelectFollowerByIDs(followerID int64, followedID int64) (domain.Follower, pkg.Error) {
	var result domain.Follower

	query := "select follower_id, followed_id, created_at from follower where follower_id = ? and followed_id = ?"
	row := f.rdb.QueryRow(query, followerID, followedID)
	err := row.Scan(&result.FollowerID, &result.FollowedID, &result.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, pkg.NewDBNotFoundError("follower", err)
		}
		return result, pkg.NewDBScanFatalError("follower", err)
	}

	return result, nil
}
