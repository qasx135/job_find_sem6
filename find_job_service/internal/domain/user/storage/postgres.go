package storageuser

import (
	"context"
	"errors"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	user "job_finder_service/internal/domain/user/model"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Storage struct {
	client Client
}

func NewStorageUser(client Client) *Storage {
	return &Storage{
		client: client,
	}
}

func (s *Storage) All(ctx context.Context) ([]user.User, error) {
	from := sqlbuilder.Select("id", "email", "password").From("users").String()

	rows, err := s.client.Query(ctx, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := make([]user.User, 0)
	for rows.Next() {
		u := user.User{}
		rows.Scan(&u.ID, &u.Email, &u.Password)
		list = append(list, u)
	}
	return list, nil
}

func (s *Storage) Create(ctx context.Context, user *user.User) error {
	q := `INSERT INTO users (email, password) VALUES ($1, $2)`
	if _, err := s.client.Exec(ctx, q, user.Email, user.Password); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil
			}
		}
	}
	return nil
}
