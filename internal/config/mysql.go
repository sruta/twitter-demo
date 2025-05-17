package config

import "os"

type MySQL struct {
	Host         string
	Port         string
	Database     string
	User         string
	Pass         string
	Driver       string
	MaxOpenConns int
	MaxIdleConns int
}

var MySQLProd = MySQL{
	Host:         os.Getenv("DB_HOST"),
	Port:         os.Getenv("DB_PORT"),
	Database:     os.Getenv("DB_NAME"),
	User:         os.Getenv("DB_USER"),
	Pass:         os.Getenv("DB_PASSWORD"),
	Driver:       "mysql",
	MaxOpenConns: 10,
	MaxIdleConns: 10,
}
