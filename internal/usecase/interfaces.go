package usecase

import (
	"context"

	"github.com/vicdevcode/schedule/internal/entity"
	"github.com/vicdevcode/schedule/internal/model"
)

type (
	// WeeklyEvent
	WeeklyEvent interface {
		FindAll(context.Context) ([]entity.WeeklyEvent, error)
		Create(context.Context, model.CreateWeeklyEvent) (*entity.WeeklyEvent, error)
		Delete(context.Context, string) error
	}
	WeeklyEventRepo interface {
		FindAll(context.Context) ([]entity.WeeklyEvent, error)
		Create(context.Context, model.CreateWeeklyEvent) (*entity.WeeklyEvent, error)
		Delete(context.Context, string) error
	}
)
