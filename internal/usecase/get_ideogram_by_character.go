package usecase

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf8"

	"mandarin-app/internal/entity"
)

type GetIdeogramByCharacterUseCase struct {
	repo entity.IdeogramRepository
}

func NewGetIdeogramByCharacterUseCase(
	repo entity.IdeogramRepository,
) *GetIdeogramByCharacterUseCase {
	return &GetIdeogramByCharacterUseCase{
		repo: repo,
	}
}

func (uc *GetIdeogramByCharacterUseCase) Execute(
	ctx context.Context,
	character string,
) (*entity.Ideogram, error) {
	character = strings.TrimSpace(character)

	if character == "" {
		return nil, entity.ErrEmptyCharacter
	}

	if utf8.RuneCountInString(character) != 1 {
		return nil, entity.ErrInvalidCharacter
	}

	ideogram, err := uc.repo.FindByCharacter(ctx, character)
	if err != nil {
		return nil, fmt.Errorf(
			"buscar ideograma %q: %w",
			character,
			err,
		)
	}

	return ideogram, nil
}
