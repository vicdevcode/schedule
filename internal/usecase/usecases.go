package usecase

import (
	"github.com/vicdevcode/init_template/internal/usecase/repo"
	"github.com/vicdevcode/init_template/pkg/config"
	"github.com/vicdevcode/init_template/pkg/postgres"
)

type UseCases struct {
	WeeklyEventUseCase WeeklyEvent
}

func New(cfg *config.Config, db *postgres.Postgres) UseCases {
	return UseCases{
		WeeklyEventUseCase: NewWeeklyEvent(repo.NewWeeklyEvent(db), cfg.ContextTimeout),
	}
}
