package repository

import "database/sql"

type IWebsRepository interface{}

type WebsRepository struct {
	db *sql.DB
}

func NewWebsRepository() *WebsRepository {
	return &WebsRepository{
		db: nil,
	}
}
