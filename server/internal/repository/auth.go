package repository

import "database/sql"

type IAuthDB interface {
	GenerateJWTToken(email, password string) (string, error)
	CheckPass(email, password string) bool
}

type AuthDB struct {
	db *sql.DB
}

func NewAuthDB(config string) *AuthDB {
	db, err := sql.Open("postgres", config)
	return &AuthDB{db: nil}
}

func (a *AuthDB) GenerateJWTToken(email, password string) (string, error) {

}

func (a *AuthDB) CheckPass(email, password string) bool {

}