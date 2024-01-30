package usecase

import (
	"context"
	"time"

	"github.com/vicdevcode/init_template/internal/entity"
	"github.com/vicdevcode/init_template/internal/model"
)

type WeeklyEventUseCase struct {
	repo       WeeklyEventRepo
	ctxTimeout time.Duration
}

func NewWeeklyEvent(r WeeklyEventRepo, t time.Duration) *WeeklyEventUseCase {
	return &WeeklyEventUseCase{
		repo:       r,
		ctxTimeout: t,
	}
}

func (uc *WeeklyEventUseCase) Create(c context.Context, data model.CreateWeeklyEvent) (*entity.WeeklyEvent, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	weeklyEvent, err := uc.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return weeklyEvent, nil
}

func (uc *WeeklyEventUseCase) FindAll(c context.Context) ([]entity.WeeklyEvent, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	weeklyEvents, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return weeklyEvents, nil
}

func (uc *WeeklyEventUseCase) Delete(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	if err := uc.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
