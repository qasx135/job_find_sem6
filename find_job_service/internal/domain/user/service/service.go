package serviceuser

import (
	"context"
	user "job_finder_service/internal/domain/user/model"
)

type PostgreRepo interface {
	All(ctx context.Context) ([]user.User, error)
	Create(ctx context.Context, employer *user.User) error
}

type Service struct {
	repo PostgreRepo
}

func NewServiceUser(repo PostgreRepo) *Service {
	return &Service{repo: repo}
}
func (s *Service) All(ctx context.Context) ([]user.User, error) {
	return s.repo.All(ctx)
}

func (s *Service) Create(ctx context.Context, employer *user.User) error {
	return s.repo.Create(ctx, employer)
}
