package service

import (
	"github.com/cucumberjaye/gophermart/internal/app/handler"
	"github.com/cucumberjaye/gophermart/internal/app/models"
)

func (s *MartService) GetBalance(userID string) (models.Balance, error) {
	balance, err := s.repository.GetBalance(userID)
	if err != nil {
		return balance, err
	}

	balance.Current -= balance.Withdrawn
	return balance, nil
}

func (s *MartService) Withdraw(userID string, withdraw models.Withdraw) error {
	balance, err := s.repository.GetBalance(userID)
	if err != nil {
		return err
	}

	if balance.Current+balance.Withdrawn-withdraw.Sum < 0 {
		return handler.ErrInsufficientFunds
	}

	withdraw.Sum = -withdraw.Sum
	return s.repository.Withdraw(userID, withdraw)
}

func (s *MartService) GetWithdrawals(userID string) ([]models.Withdraw, error) {
	return s.repository.GetWithdrawals(userID)
}
