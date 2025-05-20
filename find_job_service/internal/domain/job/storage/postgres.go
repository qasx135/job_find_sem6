package storagejob

import (
	"context"
	"errors"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	job "job_finder_service/internal/domain/job/model"
	"log/slog"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Storage struct {
	client Client
}

func NewStorageJob(client Client) *Storage {
	return &Storage{
		client: client,
	}
}

func (s *Storage) All(ctx context.Context) ([]job.Job, error) {
	from := sqlbuilder.Select(
		"id",
		"salary",
		"header",
		"experience",
		"employment",
		"schedule",
		"work_format",
		"working_hours",
		"description",
		"user_id").From("jobs").String()

	rows, err := s.client.Query(ctx, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := make([]job.Job, 0)
	for rows.Next() {
		j := job.Job{}
		rows.Scan(&j.Id, &j.Header, &j.Salary, &j.Experience, &j.Employment,
			&j.Schedule, &j.WorkFormat, &j.WorkingHours, &j.Description,
			&j.UserID)
		list = append(list, j)
	}
	slog.Info("AllJobs ------------------>", list)
	return list, nil
}

func (s *Storage) Create(ctx context.Context, job *job.Job) error {
	q := `INSERT INTO jobs (header,
	salary,
	experience,
	employment,
	schedule,
	work_format,
	working_hours,
	description,
	user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	if _, err := s.client.Exec(ctx, q,
		job.Header,
		job.Salary,
		job.Experience,
		job.Employment,
		job.Schedule,
		job.WorkFormat,
		job.WorkingHours,
		job.Description,
		job.UserID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil
			}
		}
		slog.Error("Error creating job", job, err)
	}
	slog.Info("Job created", job.Id)
	return nil
}

func (s *Storage) GetJobByID(ctx context.Context, id string) (job.Job, error) {
	query := `SELECT id, header, salary, experience, employment, schedule, work_format, working_hours, description, user_id FROM jobs WHERE id = $1`
	rows, err := s.client.Query(ctx, query, id)
	if err != nil {
		slog.Error("Error getting job by id form postgre", id, err)
	}
	defer rows.Close()
	j := job.Job{}
	for rows.Next() {
		rows.Scan(&j.Id, &j.Header, &j.Salary, &j.Experience, &j.Employment, &j.Schedule, &j.WorkFormat, &j.WorkingHours, &j.Description, &j.UserID)

	}
	slog.Info("The job found", j)
	return j, nil
}

func (s *Storage) AllJobsByUser(ctx context.Context, userID int) ([]job.Job, error) {
	query := `SELECT id,
header,
salary,
experience,
employment,
schedule,
work_format,
working_hours,
description,
user_id FROM jobs WHERE user_id = $1`
	rows, err := s.client.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := make([]job.Job, 0)
	for rows.Next() {
		j := job.Job{}
		err = rows.Scan(&j.Id, &j.Header, &j.Salary, &j.Experience, &j.Employment, &j.Schedule, &j.WorkFormat, &j.WorkingHours, &j.Description, &j.UserID)
		if err != nil {
			return nil, err
		}
		list = append(list, j)
	}
	return list, nil
}
