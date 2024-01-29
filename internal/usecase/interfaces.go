package usecase

import (
	"context"

	"github.com/vicdevcode/init_template/internal/entity"
)

type (
	// Example
	Example interface {
		GetExamples(context.Context) ([]entity.Example, error)
		CreateExample(context.Context, string) (*entity.Example, error)
	}
	ExampleRepo interface {
		GetExamples(context.Context) ([]entity.Example, error)
		CreateExample(context.Context, string) (*entity.Example, error)
	}
)
