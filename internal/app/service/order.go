package service

import (
	"time"

	"github.com/cucumberjaye/gophermart/internal/app/models"
)

func (s *MartService) SetOrder(order models.Order) error {
	order.Status = models.New
	order.UploadedAt = time.Now()

	return s.repository.SetOrder(order)
}

func (s *MartService) GetOrders(userID string) ([]models.Order, error) {
	return s.repository.GetOrders(userID)
}
