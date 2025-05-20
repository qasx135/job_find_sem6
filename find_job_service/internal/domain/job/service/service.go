package servicejob

import (
	"context"
	job "job_finder_service/internal/domain/job/model"
)

type PostgreRepo interface {
	All(ctx context.Context) ([]job.Job, error)
	Create(ctx context.Context, employer *job.Job) error
	GetJobByID(ctx context.Context, id string) (job.Job, error)
	AllJobsByUser(ctx context.Context, userID int) ([]job.Job, error)
}

type Service struct {
	repo PostgreRepo
}

func NewServiceJob(repo PostgreRepo) *Service {
	return &Service{repo: repo}
}
func (s *Service) All(ctx context.Context) ([]job.Job, error) {
	return s.repo.All(ctx)
}

func (s *Service) Create(ctx context.Context, employer *job.Job) error {
	return s.repo.Create(ctx, employer)
}

func (s *Service) GetJobByID(ctx context.Context, id string) (job.Job, error) {
	return s.repo.GetJobByID(ctx, id)
}

func (s *Service) AllJobsByUser(ctx context.Context, userID int) ([]job.Job, error) {
	return s.repo.AllJobsByUser(ctx, userID)
}
