package repo

import (
	"context"
	"fmt"

	"github.com/vicdevcode/init_template/internal/entity"
	"github.com/vicdevcode/init_template/pkg/postgres"
)

type ExampleRepo struct {
	*postgres.Postgres
}

func NewExample(db *postgres.Postgres) *ExampleRepo {
	return &ExampleRepo{db}
}

func (r *ExampleRepo) CreateExample(ctx context.Context, message string) (*entity.Example, error) {
	example := &entity.Example{Message: message}
	result := r.WithContext(ctx).Create(example)
	if result.Error != nil {
		return nil, fmt.Errorf("ExampleRepo - CreateExample: %w", result.Error)
	}
	return example, nil
}

func (r *ExampleRepo) GetExamples(ctx context.Context) ([]entity.Example, error) {
	examples := []entity.Example{}
	result := r.WithContext(ctx).Find(&examples)
	if result.Error != nil {
		return nil, fmt.Errorf("ExampleRepo - GetExamples: %w", result.Error)
	}
	return examples, nil
}
