package service

import (
	"context"

	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/entity"
	port "github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/ports/repositories"
)

type UserService struct {
	userRepo port.UserRepository
}

//var _ ports.UserRepository = (*UserRepository)(nil)

func NewUserService(userRepo port.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	return us.userRepo.CreateUser(ctx, user)
}

func (us *UserService) DeleteUser(ctx context.Context, userID string) error {
	return us.userRepo.DeleteUser(ctx, userID)
}

func (us *UserService) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	return us.userRepo.GetUser(ctx, userID)
}

func (us *UserService) UpdateUser(ctx context.Context, user entity.User) error {
	return us.userRepo.UpdateUser(ctx, user)
}
