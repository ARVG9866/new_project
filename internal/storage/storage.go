package storage

import "github.com/jmoiron/sqlx"

type IStorage interface {
}

type storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) IStorage {
	return &storage{db: db}
}
