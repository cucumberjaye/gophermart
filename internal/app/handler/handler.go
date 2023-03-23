package handler

import (
	"errors"

	"github.com/cucumberjaye/gophermart/internal/app/middleware"
	"github.com/cucumberjaye/gophermart/internal/app/models"

	"github.com/go-chi/chi/v5"
)

var (
	ErrInvalidOrder      = errors.New("invalid order")
	ErrBadOrder          = errors.New("bad order")
	ErrUserOrderExists   = errors.New("order was set this user")
	ErrOrderExists       = errors.New("order was set other user")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type AuthService interface {
	CreateUser(user models.RegisterUser) error
	GenerateToken(user models.LoginUser) (string, error)
}

type OrderService interface {
	SetOrder(order models.Order) error
	GetOrders(userID string) ([]models.Order, error)
}

type BalanceService interface {
	GetBalance(userID string) (models.Balance, error)
	Withdraw(userID string, withdraw models.Withdraw) error
	GetWithdrawals(userID string) ([]models.Withdraw, error)
}

type MartService interface {
	AuthService
	OrderService
	BalanceService
}

type Handler struct {
	service MartService
}

func New(service MartService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.GzipCompress, middleware.GzipDecompress)

	r.Route("/user", func(r chi.Router) {
		r.Post("/register", h.register)
		r.Post("/login", h.login)

		r.With(middleware.Authentication).Group(func(r chi.Router) {
			r.Post("/orders", h.setOrder)
			r.Get("/orders", h.getOrders)

			r.Get("/balance", h.getBalance)
			r.Post("/balance/withdraw", h.withdraw)

			r.Get("/withdrawals", h.getWithdraws)
		})
	})

	return r
}
