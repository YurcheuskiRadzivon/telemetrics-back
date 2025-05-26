package repositories

import (
	"context"

	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/core/entity"
	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/database/queries"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ViewOptRepository struct {
	queries *queries.Queries
	pool    *pgxpool.Pool
}

//var _ ports.UserRepository = (*UserRepository)(nil)

func NewViewOptRepository(db *queries.Queries, pool *pgxpool.Pool) *ViewOptRepository {
	return &ViewOptRepository{
		queries: db,
		pool:    pool,
	}
}

func (vo *ViewOptRepository) CreateViewOptions(ctx context.Context, viewOpt *entity.ViewOptions) error {
	params := queries.CreateViewOptionsParams{
		UserID:            viewOpt.UserID,
		ChannelCount:      viewOpt.ChannelCount,
		Tittle:            viewOpt.Tittle,
		About:             viewOpt.About,
		ChannelID:         viewOpt.ChannelID,
		ChannelDate:       viewOpt.ChannelDate,
		ParticipantsCount: viewOpt.ParticipantsCount,
		Photo:             viewOpt.Photo,
		MessageCount:      viewOpt.MessageCount,
		MessageID:         viewOpt.MessageID,
		Views:             viewOpt.Views,
		PostDate:          viewOpt.PostDate,
		ReactionsCount:    viewOpt.ReactionsCount,
		Reactions:         viewOpt.Reactions,
	}
	return vo.queries.CreateViewOptions(ctx, params)
}

func (vo *ViewOptRepository) DeleteViewOptions(ctx context.Context, userID string) error {
	return vo.DeleteViewOptions(ctx, userID)
}

func (vo *ViewOptRepository) GetViewOptions(ctx context.Context, userID string) (*entity.ViewOptions, error) {
	viewOpt, err := vo.queries.GetViewOptions(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &entity.ViewOptions{
		UserID:            viewOpt.UserID,
		ChannelCount:      viewOpt.ChannelCount,
		Tittle:            viewOpt.Tittle,
		About:             viewOpt.About,
		ChannelID:         viewOpt.ChannelID,
		ChannelDate:       viewOpt.ChannelDate,
		ParticipantsCount: viewOpt.ParticipantsCount,
		Photo:             viewOpt.Photo,
		MessageCount:      viewOpt.MessageCount,
		MessageID:         viewOpt.MessageID,
		Views:             viewOpt.Views,
		PostDate:          viewOpt.PostDate,
		ReactionsCount:    viewOpt.ReactionsCount,
		Reactions:         viewOpt.Reactions,
	}, nil
}

func (vo *ViewOptRepository) UpdateViewOptions(ctx context.Context, viewOpt entity.ViewOptions) error {
	return vo.queries.UpdateViewOptions(ctx, queries.UpdateViewOptionsParams{
		UserID:            viewOpt.UserID,
		ChannelCount:      viewOpt.ChannelCount,
		Tittle:            viewOpt.Tittle,
		About:             viewOpt.About,
		ChannelID:         viewOpt.ChannelID,
		ChannelDate:       viewOpt.ChannelDate,
		ParticipantsCount: viewOpt.ParticipantsCount,
		Photo:             viewOpt.Photo,
		MessageCount:      viewOpt.MessageCount,
		MessageID:         viewOpt.MessageID,
		Views:             viewOpt.Views,
		PostDate:          viewOpt.PostDate,
		ReactionsCount:    viewOpt.ReactionsCount,
		Reactions:         viewOpt.Reactions,
	})
}
