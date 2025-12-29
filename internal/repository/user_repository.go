package repository

import (
	"context"
	"database/sql"
	"errors"
	"goserver/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(ctx context.Context, tx *sql.Tx, user domain.User) error {
	query := `
	INSERT INTO users (email, password)	
	VALUES ($1, $2)
	RETURNING id
	`

	return tx.QueryRowContext(ctx, query, user.Email, user.Password).Scan(&user.ID)
}

func (r *UserRepository) FindManyUsers(ctx context.Context, tx *sql.Tx) ([]domain.User, error) {
	query := `SELECT id, email, password FROM users`

	var users []domain.User

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, tx *sql.Tx, email string) (*domain.User, error) {
	query := `SELECT id, email, password FROM users WHERE email = $1`

	var user domain.User

	row := tx.QueryRowContext(ctx, query, email)

	if tx == nil {
		row = r.db.QueryRowContext(ctx, query, email)
	} else {
		row = tx.QueryRowContext(ctx, query, email)
	}

	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
