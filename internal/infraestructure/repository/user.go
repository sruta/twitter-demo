package repository

import (
	"database/sql"
	"errors"
	"twitter-uala/internal/domain"
	"twitter-uala/pkg"
)

type IUser interface {
	Select() ([]domain.User, pkg.Error)
	SelectByID(id int64) (domain.User, pkg.Error)
	SelectByEmail(email string) (domain.User, pkg.Error)
	SelectByUsername(username string) (domain.User, pkg.Error)
	Insert(user domain.User) (domain.User, pkg.Error)
	UpdateByID(user domain.User) (domain.User, pkg.Error)
}

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

	row := u.rdb.QueryRow("select id, email, password, username from user where email = ?", email)
	err := row.Scan(&result.ID, &result.Email, &result.Password, &result.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, pkg.NewDBNotFoundError("user", err)
		}
		return result, pkg.NewDBScanFatalError("user", err)
	}

	return result, nil
}

func (u User) SelectByUsername(username string) (domain.User, pkg.Error) {
	var result domain.User

	row := u.rdb.QueryRow("select id, email, username from user where username = ?", username)
	err := row.Scan(&result.ID, &result.Email, &result.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, pkg.NewDBNotFoundError("user", err)
		}
		return result, pkg.NewDBScanFatalError("user", err)
	}

	return result, nil
}

func (u User) Insert(user domain.User) (domain.User, pkg.Error) {
	query := "insert into user(email, password, username) values (?,?, ?)"
	result, err := u.rdb.Exec(query, user.Email, user.Password, user.Username)
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

func (u User) UpdateByID(user domain.User) (domain.User, pkg.Error) {
	var err error
	query := "update user set username = ? where id = ?"
	_, err = u.rdb.Exec(query, user.Username, user.ID)
	if err != nil {
		return user, pkg.NewDBFatalError("update user in", err)
	}

	return user, nil
}
