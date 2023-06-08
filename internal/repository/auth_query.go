package repository

import (
	"casbin-go_gin/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Auth struct {
	db *pgxpool.Pool
}

func NewAuth(db *pgxpool.Pool) *Auth {
	return &Auth{db: db}
}

func (r *Auth) Crete(user models.User) error {
	query := fmt.Sprintf(`INSERT INTO %s (name, password, phone, address, role) VALUES ($1, $2, $3, $4, $5)`, "users")
	_, err := r.db.Exec(context.Background(), query, user.Name, user.Password, user.Phone, user.Address, user.Role)
	if err != nil {
		return err
	}
	return nil
}

func (r *Auth) Get(name, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf(`SELECT id, name, phone, address, role FROM %s WHERE name = $1 AND password = $2 `, "users")
	if err := r.db.QueryRow(context.Background(), query, name, password).Scan(&user.Id, &user.Name, &user.Phone, &user.Address, &user.Role); err != nil {
		return models.User{}, err
	}
	return user, nil
}
