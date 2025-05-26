package repositories

import (
	"context"

	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/entity"
)

type ViewOptionRepository interface {
	CreateViewOptions(ctx context.Context, viewOpt *entity.ViewOptions) error
	DeleteViewOptions(ctx context.Context, userID string) error
	GetViewOptions(ctx context.Context, userID string) (*entity.ViewOptions, error)
	UpdateViewOptions(ctx context.Context, viewOpt entity.ViewOptions) error
}
