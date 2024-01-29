package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/vicdevcode/init_template/internal/entity"
)

type ExampleUseCase struct {
	repo       ExampleRepo
	ctxTimeout time.Duration
}

func NewExample(r ExampleRepo, t time.Duration) *ExampleUseCase {
	return &ExampleUseCase{
		repo:       r,
		ctxTimeout: t,
	}
}

func (uc *ExampleUseCase) GetExamples(c context.Context) ([]entity.Example, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	examples, err := uc.repo.GetExamples(ctx)
	if err != nil {
		return nil, fmt.Errorf("ExampleUseCase - GetExamples - repo.GetExamples: %w", err)
	}
	return examples, nil
}

func (uc *ExampleUseCase) CreateExample(
	c context.Context,
	message string,
) (*entity.Example, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	example, err := uc.repo.CreateExample(ctx, message)
	if err != nil {
		return nil, fmt.Errorf("ExampleUseCase - CreateExample - repo.CreateExample: %w", err)
	}
	return example, nil
}
