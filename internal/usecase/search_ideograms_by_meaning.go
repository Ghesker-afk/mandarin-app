package usecase

import (
	"context"
	"fmt"
	"strings"

	"mandarin-app/internal/entity"
)

type SearchIdeogramsByMeaningUseCase struct {
	repo entity.IdeogramRepository
}

func NewSearchIdeogramsByMeaningUseCase(
	repo entity.IdeogramRepository,
) *SearchIdeogramsByMeaningUseCase {
	return &SearchIdeogramsByMeaningUseCase{
		repo: repo,
	}
}

func (uc *SearchIdeogramsByMeaningUseCase) Execute(
	ctx context.Context,
	meaning string,
) ([]entity.Ideogram, error) {
	meaning = strings.TrimSpace(meaning)

	if meaning == "" {
		return nil, entity.ErrEmptyMeaning
	}

	ideograms, err := uc.repo.SearchByMeaning(ctx, meaning)
	if err != nil {
		return nil, fmt.Errorf(
			"buscar ideogramas pelo significado %q: %w",
			meaning,
			err,
		)
	}

	return ideograms, nil
}
