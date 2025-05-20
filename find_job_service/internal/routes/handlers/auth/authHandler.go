package auth

import (
	"context"
	"fmt"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	user "job_finder_service/internal/domain/user/model"
	"log/slog"
	"net/http"
	"time"
)

var jwtKey = []byte("dont_look_my_key")

type AuthService interface {
	Register(ctx context.Context, email string, password []byte) (*user.User, error)
	Login(ctx context.Context, email, password string) (*user.User, error)
}

type AuthHandler struct {
	Service AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler {

	return &AuthHandler{
		Service: service,
	}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	req := &RegisterRequest{}
	if err := render.DecodeJSON(r.Body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Error hashing password", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := &user.User{
		Email:    req.Email,
		Password: hash,
	}
	user, err = h.Service.Register(r.Context(), req.Email, hash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error registering user", err)
		return
	}
	expireTime := time.Now().Add(time.Hour * 24)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expireTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", user.ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error signing token", err)
		return
	}
	render.JSON(w, r, &TokenResponse{
		Token: tokenString,
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := &LoginRequest{}
	if err := render.DecodeJSON(r.Body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("Error decoding body", err)
		return
	}
	user, err := h.Service.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		slog.Error("Error authenticating user", err)
		return
	}

	expireTime := time.Now().Add(time.Hour * 24)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expireTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", user.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error signing token", err)
		return
	}
	render.JSON(w, r, &TokenResponse{
		Token: tokenString,
	})
}
