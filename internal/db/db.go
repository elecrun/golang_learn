package db

import (
	"database/sql"
	"highqps-cache/internal/model"
)

type DB struct {
	conn *sql.DB
}

func NewDB(conn *sql.DB) *DB {
	return &DB{conn: conn}
}

func (d *DB) GetByID(id int64) (*model.Item, error) {
	row := d.conn.QueryRow("SELECT id, name, value FROM items WHERE id = ?", id)

	item := &model.Item{}
	err := row.Scan(&item.ID, &item.Name, &item.Value)
	if err != nil {
		return nil, err
	}

	return item, nil
}