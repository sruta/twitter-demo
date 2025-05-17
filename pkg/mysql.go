package pkg

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"twitter-uala/internal/configs"
)

type DB interface {
	WithTransaction(func(*sql.Tx) Error) Error
	Close() Error
}

type MySQL struct {
	*sql.DB
}

func NewMySQL(c configs.MySQL) (*MySQL, Error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", c.User, c.Pass, c.Host, c.Port, c.Database)

	db, err := sql.Open(c.Driver, connectionString)
	if err != nil {
		return nil, NewFatalError("could not connect to db", err)
	}
	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetConnMaxLifetime(10 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, NewFatalError("could not ping db", err)
	}

	return &MySQL{db}, nil
}

func (db MySQL) WithTransaction(txFunc func(tx *sql.Tx) Error) Error {
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return NewFatalError("could not begin db tx", err)
	}

	var txErr Error
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Rollbacking due panic ")
			err = tx.Rollback()
			panic(p)
		} else if txErr != nil {
			fmt.Printf("Rollbacking due txFunc error ")
			_ = tx.Rollback() // err is non-nil, don't change it
		} else {
			fmt.Printf("Commiting txFunc ")
			err = tx.Commit() // err is nil, if Commit returns Error update err
		}
	}()

	txErr = txFunc(tx)
	return txErr
}

func (db MySQL) Close() Error {
	err := db.DB.Close()
	if err != nil {
		return NewFatalError("could not close db connection", err)
	}
	return nil
}

type RowScanner interface {
	Scan(...interface{}) error
}
