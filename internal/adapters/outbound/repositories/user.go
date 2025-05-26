package repositories

import (
	"context"

	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/entity"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/database/queries"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	queries *queries.Queries
	pool    *pgxpool.Pool
}

//var _ ports.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *queries.Queries, pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		queries: db,
		pool:    pool,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	params := queries.CreateUserParams{
		UserID:      user.UserID,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
	}

	return r.queries.CreateUser(ctx, params)
}

func (r *UserRepository) DeleteUser(ctx context.Context, userID string) error {
	return r.DeleteUser(ctx, userID)
}

func (r *UserRepository) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	user, err := r.queries.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		UserID:      user.UserID,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
	}, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user entity.User) error {
	return r.queries.UpdateUser(ctx, queries.UpdateUserParams{
		UserID:      user.UserID,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
	})
}
