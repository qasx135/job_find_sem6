package authservice

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	user "job_finder_service/internal/domain/user/model"
	"log/slog"
)

type AuthService struct {
	pool *pgxpool.Pool
}

func NewAuthService(pool *pgxpool.Pool) *AuthService {
	return &AuthService{pool: pool}
}

func (s *AuthService) Register(ctx context.Context, email string, password []byte) (*user.User, error) {
	_, err := s.pool.Exec(
		ctx,
		"INSERT INTO users (email, password) VALUES ($1, $2)",
		email, password,
	)
	var user user.User

	err = s.pool.QueryRow(ctx, "SELECT id, email, password FROM users WHERE email = $1", email).Scan(
		&user.ID, &user.Email, &user.Password,
	)
	if err != nil {
		slog.Error("Error registering user", err)
		return nil, err
	}

	return &user, err
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*user.User, error) {
	var user user.User

	err := s.pool.QueryRow(ctx, "SELECT id, email, password FROM users WHERE email = $1", email).Scan(
		&user.ID, &user.Email, &user.Password,
	)

	if err != nil {
		return nil, errors.New("didnt find user from service auth")
	}

	slog.Info("User found", "email", user.Email, "hash", user.Password)

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		slog.Error("Password mismatch", "input", password, "stored_hash", user.Password)
		return nil, errors.New("invalid credentials!!!")
	}

	return &user, nil
}
