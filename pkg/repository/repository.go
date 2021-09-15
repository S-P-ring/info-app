package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type InfoList interface {
}

type InfoListItem interface {
}

type Repository struct {
	Authorization
	InfoList
	InfoListItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
