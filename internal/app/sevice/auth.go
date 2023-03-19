package service

import (
	"github.com/cucumberjaye/gophermart/configs"
	"github.com/cucumberjaye/gophermart/internal/app/models"
	"github.com/cucumberjaye/gophermart/pkg/hasher"
	"github.com/cucumberjaye/gophermart/pkg/token"
	"github.com/google/uuid"
)

func (s *MartService) CreateUser(user models.RegisterUser) error {
	user.Password = hasher.GeneratePasswordHash(user.Password)
	id := uuid.New().String()
	return s.repository.CreateUser(id, user)
}

func (s *MartService) GenerateToken(loginUser models.LoginUser) (string, error) {
	loginUser.Password = hasher.GeneratePasswordHash(loginUser.Password)
	user, err := s.repository.GetUser(loginUser)
	if err != nil {
		return "", err
	}

	return token.GenerateToken(user.ID, []byte(configs.SigningKey))
}
