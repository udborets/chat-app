package websocket

import "database/sql"

type IWsRepository interface{}

type WsRepository struct {
	db *sql.DB
}
