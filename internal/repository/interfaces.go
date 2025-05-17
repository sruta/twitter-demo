package repository

import (
	"database/sql"
	"twitter-uala/internal/domain"
	"twitter-uala/pkg"
)

type IUser interface {
	Select() ([]domain.User, pkg.Error)
	SelectByID(id int64) (domain.User, pkg.Error)
	SelectByEmail(email string) (domain.User, pkg.Error)
	SelectByEmailAndPassword(email string, password string) (domain.User, pkg.Error)
	Insert(tx *sql.Tx, user domain.User) (domain.User, pkg.Error)
}
