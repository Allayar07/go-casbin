package repository

import "github.com/jackc/pgx/v4/pgxpool"

type Repository struct {
	Auth *Auth
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{Auth: NewAuth(db)}
}
