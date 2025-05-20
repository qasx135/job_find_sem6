package handlers

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	job "job_finder_service/internal/domain/job/model"
	resume "job_finder_service/internal/domain/resume/model"
	user "job_finder_service/internal/domain/user/model"
	"job_finder_service/internal/middleware"
	"job_finder_service/internal/routes/handlers/auth"
	"log/slog"
	"net/http"
	"strconv"
)

type ServiceJob interface {
	Create(ctx context.Context, job *job.Job) error
	All(ctx context.Context) ([]job.Job, error)
	GetJobByID(ctx context.Context, id string) (job.Job, error)
	AllJobsByUser(ctx context.Context, userID int) ([]job.Job, error)
}

type ServiceResume interface {
	Create(ctx context.Context, resume *resume.Resume) error
	All(ctx context.Context) ([]resume.Resume, error)
	AllByUser(ctx context.Context, userID int) ([]resume.Resume, error)
}

type ServiceUser interface {
	Create(ctx context.Context, user *user.User) error
	All(ctx context.Context) ([]user.User, error)
}

type Handler struct {
	AuthHandler *auth.AuthHandler
	ServiceJob  ServiceJob
	ServiceRes  ServiceResume
	ServiceUser ServiceUser
	ctx         context.Context
}

func NewHandler(ctx context.Context,
	jobService ServiceJob,
	resumeService ServiceResume,
	userService ServiceUser,
	authHandler *auth.AuthHandler) *Handler {
	return &Handler{
		ServiceJob:  jobService,
		ServiceRes:  resumeService,
		ServiceUser: userService,
		AuthHandler: authHandler,
		ctx:         ctx,
	}
}

func (s *Handler) CreateJob(w http.ResponseWriter, r *http.Request) {
	job := &job.Job{}
	if err := render.DecodeJSON(r.Body, job); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("Error decoding job json", err)
	}
	// Получаем userID из контекста
	userIDStr := r.Context().Value(middleware.UserIDKey).(string)
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("Invalid user ID", err)
		return
	}

	// Устанавливаем user_id для резюме
	job.UserID = userID
	err = s.ServiceJob.Create(s.ctx, job)
	slog.Info("Job created handler", job.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error creating job", err)
	}
}

func (s *Handler) GetJobById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	slog.Info("GetJobById handler", id)
	job, err := s.ServiceJob.GetJobByID(s.ctx, id)
	fmt.Println(job)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error getting job from handler", err)
		return
	}
	render.JSON(w, r, job)
}
func (s *Handler) AllJobs(w http.ResponseWriter, r *http.Request) {
	allJobs, err := s.ServiceJob.All(s.ctx)
	slog.Info("AllJobs ------------------>", allJobs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error getting all jobs from handlers", err)
	}
	render.JSON(w, r, allJobs)
}

func (s *Handler) AllJobsByUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value(middleware.UserIDKey).(string)
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("Invalid user ID", err)
		return
	}
	allJobs, err := s.ServiceJob.AllJobsByUser(s.ctx, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error getting all jobs from handler", err)
		return
	}
	render.JSON(w, r, allJobs)
}

func (s *Handler) CreateResume(w http.ResponseWriter, r *http.Request) {
	resume := &resume.Resume{}
	if err := render.DecodeJSON(r.Body, resume); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("Error decoding resume json", err)
		return
	}

	// Получаем userID из контекста
	userIDStr := r.Context().Value(middleware.UserIDKey).(string)
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("Invalid user ID", err)
		return
	}

	// Устанавливаем user_id для резюме
	resume.UserID = userID

	// Передаем в сервис
	err = s.ServiceRes.Create(r.Context(), resume)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error creating resume", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (s *Handler) AllResume(w http.ResponseWriter, r *http.Request) {
	allResume, err := s.ServiceRes.All(s.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error getting all resume from handlers", err)
	}
	render.JSON(w, r, allResume)
}

func (s *Handler) AllResumeByUser(w http.ResponseWriter, r *http.Request) {
	// Получаем userID из контекста
	userIDStr := r.Context().Value(middleware.UserIDKey).(string)
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("Invalid user ID", err)
		return
	}

	// Получаем только свои резюме
	allResume, err := s.ServiceRes.AllByUser(r.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error getting all resume from handlers", err)
		return
	}

	render.JSON(w, r, allResume)
}

func (s *Handler) AllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers, err := s.ServiceUser.All(s.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error getting all users from handlers", err)
	}
	render.JSON(w, r, allUsers)
}

//------------------------------------auth-------------------------------------

func (s *Handler) HandleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := render.DecodeJSON(r.Body, req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			slog.Error("Error decoding json", err)
		}
	}
}
