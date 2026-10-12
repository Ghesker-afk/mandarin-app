package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"mandarin-app/internal/entity"
)

type CreateIdeogramUseCase struct {
	ideogramRepo entity.IdeogramRepository
}

// Constructor
func NewCreateIdeogramUseCase(
	repo entity.IdeogramRepository,
) *CreateIdeogramUseCase {
	return &CreateIdeogramUseCase{
		ideogramRepo: repo,
	}
}

func (uc *CreateIdeogramUseCase) Execute(
	ctx context.Context,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	if ideog == nil {
		return nil, errors.New("ideogram não pode ser nil")
	}

	ideog.Character = strings.TrimSpace(ideog.Character)
	ideog.Pinyin = strings.TrimSpace(ideog.Pinyin)
	ideog.Radical = strings.TrimSpace(ideog.Radical)

	if ideog.Character == "" {
		return nil, entity.ErrEmptyCharacter
	}

	if ideog.Pinyin == "" {
		return nil, errors.New("pinyin é obrigatório")
	}

	meanings := make([]string, 0, len(ideog.Meanings))

	for _, meaning := range ideog.Meanings {
		meaning = strings.TrimSpace(meaning)

		if meaning != "" {
			meanings = append(meanings, meaning)
		}
	}

	if len(meanings) == 0 {
		return nil, errors.New("pelo menos um significado é obrigatório")
	}

	ideog.Meanings = meanings

	created, err := uc.ideogramRepo.Create(ctx, ideog)
	if err != nil {
		if errors.Is(err, entity.ErrAlreadyExists) {
			return nil, entity.ErrAlreadyExists
		}

		return nil, fmt.Errorf("criar ideograma: %w", err)
	}

	return created, nil
}
