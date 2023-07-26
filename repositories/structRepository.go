package repositories

import "database/sql"

type UsersUserRepository struct {
	DbHandler   *sql.DB
	Transaction *sql.Tx
}

type UsersEduRepository struct {
	DbHandler   *sql.DB
	Transaction *sql.Tx
}

type UsersPhoneRepository struct {
	DbHandler   *sql.DB
	Transaction *sql.Tx
}

type UsersMediaRepository struct {
	DbHandler   *sql.DB
	Transaction *sql.Tx
}
