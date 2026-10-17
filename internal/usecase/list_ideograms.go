package usecase

import (
	"context"
	"fmt"

	"mandarin-app/internal/entity"
)

type ListIdeogramsUseCase struct {
	repo entity.IdeogramRepository
}

func NewListIdeogramsUseCase(
	repo entity.IdeogramRepository,
) *ListIdeogramsUseCase {
	return &ListIdeogramsUseCase{
		repo: repo,
	}
}

func (uc *ListIdeogramsUseCase) Execute(
	ctx context.Context,
) ([]entity.Ideogram, error) {
	ideograms, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("listar ideogramas: %w", err)
	}

	return ideograms, nil
}
