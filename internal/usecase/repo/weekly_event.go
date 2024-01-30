package repo

import (
	"context"

	"github.com/vicdevcode/schedule/internal/entity"
	"github.com/vicdevcode/schedule/internal/model"
	"github.com/vicdevcode/schedule/pkg/postgres"
)

type WeeklyEventRepo struct {
	*postgres.Postgres
}

func NewWeeklyEvent(db *postgres.Postgres) *WeeklyEventRepo {
	return &WeeklyEventRepo{db}
}

func (r *WeeklyEventRepo) Create(ctx context.Context, data model.CreateWeeklyEvent) (*entity.WeeklyEvent, error) {
	weeklyEvent := &entity.WeeklyEvent{
		Name:      data.Name,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
		WeekDay:   data.WeekDay,
		WeekType:  data.WeekType,
	}
	result := r.WithContext(ctx).Create(weeklyEvent)
	if result.Error != nil {
		return nil, result.Error
	}
	return weeklyEvent, nil
}

func (r *WeeklyEventRepo) FindAll(ctx context.Context) ([]entity.WeeklyEvent, error) {
	weeklyEvents := []entity.WeeklyEvent{}
	result := r.WithContext(ctx).Find(&weeklyEvents)
	if result.Error != nil {
		return nil, result.Error
	}
	return weeklyEvents, nil
}

func (r *WeeklyEventRepo) Delete(ctx context.Context, id string) error {
	result := r.WithContext(ctx).Where("id = ?", id).Delete(&entity.WeeklyEvent{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
