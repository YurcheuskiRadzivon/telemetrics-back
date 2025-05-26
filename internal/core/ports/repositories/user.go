package repositories

import (
	"context"

	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID string) error
	GetUser(ctx context.Context, userID string) (*entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) error
}
