package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
	middleware2 "job_finder_service/internal/middleware"
	"job_finder_service/internal/routes/handlers"
)

type Router struct {
	Router  *chi.Mux
	Handler handlers.Handler
}

func NewRouter(h *handlers.Handler, pool *pgxpool.Pool) *Router {
	cs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(cs.Handler)

	r.Post("/register", h.AuthHandler.Register)
	r.Post("/login", h.AuthHandler.Login)
	r.Get("/all-users", h.AllUsers)
	r.Get("/all-jobs-home", h.AllJobs)
	r.Get("/jobs/{id}", h.GetJobById)

	r.Group(func(r chi.Router) {
		r.Use(middleware2.AuthMiddleware(pool))

		r.Get("/all-jobs", h.AllJobsByUser)
		r.Get("/all-resumes", h.AllResumeByUser)

		r.Post("/new-job", h.CreateJob)
		r.Post("/new-resume", h.CreateResume)
	})

	return &Router{
		Router:  r,
		Handler: *h,
	}
}
