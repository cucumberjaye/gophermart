package service

import (
	"github.com/cucumberjaye/gophermart/internal/app/models"
)

type AuthRepository interface {
	GetUser(user models.LoginUser) (models.User, error)
	CreateUser(id string, user models.RegisterUser) error
}

type OrderRepository interface {
	SetOrder(order models.Order) error
	GetOrders(userID string) ([]models.Order, error)
}

type BalanceRepository interface {
	GetBalance(userID string) (models.Balance, error)
	Withdraw(userID string, withdraw models.Withdraw) error
	GetWithdrawals(userID string) ([]models.Withdraw, error)
}

type MartRepository interface {
	AuthRepository
	OrderRepository
	BalanceRepository
}

type MartService struct {
	repository MartRepository
}

func New(repository MartRepository) *MartService {
	return &MartService{
		repository: repository,
	}
}
