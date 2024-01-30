package usecase

import (
	"github.com/vicdevcode/schedule/internal/usecase/repo"
	"github.com/vicdevcode/schedule/pkg/config"
	"github.com/vicdevcode/schedule/pkg/postgres"
)

type UseCases struct {
	WeeklyEventUseCase WeeklyEvent
}

func New(cfg *config.Config, db *postgres.Postgres) UseCases {
	return UseCases{
		WeeklyEventUseCase: NewWeeklyEvent(repo.NewWeeklyEvent(db), cfg.ContextTimeout),
	}
}
