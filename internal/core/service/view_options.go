package service

import (
	"context"

	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/entity"
	port "github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/ports/repositories"
)

type ViewOptService struct {
	viewOptRepo port.ViewOptionRepository
}

//var _ ports.UserRepository = (*UserRepository)(nil)

func NewviewOptService(viewOptRepo port.ViewOptionRepository) *ViewOptService {
	return &ViewOptService{
		viewOptRepo: viewOptRepo,
	}
}

func (vo *ViewOptService) CreateViewOptions(ctx context.Context, viewOpt *entity.ViewOptions) error {
	return vo.viewOptRepo.CreateViewOptions(ctx, viewOpt)
}

func (vo *ViewOptService) DeleteViewOptions(ctx context.Context, userID string) error {
	return vo.viewOptRepo.DeleteViewOptions(ctx, userID)
}

func (vo *ViewOptService) GetViewOptions(ctx context.Context, userID string) (*entity.ViewOptions, error) {
	return vo.viewOptRepo.GetViewOptions(ctx, userID)
}

func (vo *ViewOptService) UpdateViewOptions(ctx context.Context, viewOpt entity.ViewOptions) error {
	return vo.viewOptRepo.UpdateViewOptions(ctx, viewOpt)
}
