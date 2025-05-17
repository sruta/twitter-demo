package repository

import (
	"database/sql"
	"errors"
	"twitter-uala/internal/domain"
	"twitter-uala/pkg"
)

type User struct {
	rdb *pkg.MySQL
}

func NewUser(rdb *pkg.MySQL) User {
	return User{
		rdb: rdb,
	}
}

func (u User) Select() ([]domain.User, pkg.Error) {
	result := []domain.User{}

	rows, err := u.rdb.Query("select id, email from user")
	if err != nil {
		return result, pkg.NewDBFatalError("get users from", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.ID, &user.Email)
		if err != nil {
			return result, pkg.NewDBScanFatalError("user", err)
		}

		result = append(result, user)
	}

	return result, nil
}

func (u User) SelectByID(id int64) (domain.User, pkg.Error) {
	var result domain.User

	row := u.rdb.QueryRow("select id, email, username from user where id = ?", id)
	err := row.Scan(&result.ID, &result.Email, &result.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, pkg.NewDBNotFoundError("user", err)
		}
		return result, pkg.NewDBScanFatalError("user", err)
	}

	return result, nil
}

func (u User) SelectByEmail(email string) (domain.User, pkg.Error) {
	var result domain.User

	row := u.rdb.QueryRow("select id, email, username from user where email = ?", email)
	err := row.Scan(&result.ID, &result.Email, &result.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, pkg.NewDBNotFoundError("user", err)
		}
		return result, pkg.NewDBScanFatalError("user", err)
	}

	return result, nil
}

func (u User) SelectByEmailAndPassword(email string, password string) (domain.User, pkg.Error) {
	var result domain.User

	row := u.rdb.QueryRow("select id, email, username from user where email = ? and password = ?",
		email, password)
	err := row.Scan(&result.ID, &result.Email, &result.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, pkg.NewDBNotFoundError("user", err)
		}
		return result, pkg.NewDBScanFatalError("user", err)
	}

	return result, nil
}

func (u User) Insert(tx *sql.Tx, user domain.User) (domain.User, pkg.Error) {
	var result sql.Result
	var err error
	query := "insert into user(email, password, username) values (?,?, ?)"
	if tx != nil {
		result, err = tx.Exec(query, user.Email, user.Password, user.Username)
	} else {
		result, err = u.rdb.Exec(query, user.Email, user.Password, user.Username)
	}
	if err != nil {
		return user, pkg.NewDBFatalError("insert user into", err)
	}

	user.ID, err = result.LastInsertId()
	if err != nil {
		return user, pkg.NewDBFatalError("insert user into", err)
	}
	user.Password = ""

	return user, nil
}
