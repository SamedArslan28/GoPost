package service

import (
	"context"

	"github.com/SamedArslan28/gopost/internal/models"
	"github.com/SamedArslan28/gopost/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s UserService) Register(ctx context.Context, user *models.User) (*models.User, error) {
	savedUser, err := s.repo.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return savedUser, nil
}

func (s UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
