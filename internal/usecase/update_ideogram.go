package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"mandarin-app/internal/entity"
)

type UpdateIdeogramUseCase struct {
	ideogramRepo entity.IdeogramRepository
}

func NewUpdateIdeogramUseCase(
	repo entity.IdeogramRepository,
) *UpdateIdeogramUseCase {
	return &UpdateIdeogramUseCase{
		ideogramRepo: repo,
	}
}

func (uc *UpdateIdeogramUseCase) Execute(
	ctx context.Context,
	character string,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	if ideog == nil {
		return nil, errors.New("ideogram não pode ser nil")
	}

	character = strings.TrimSpace(character)
	if character == "" {
		return nil, entity.ErrEmptyCharacter
	}

	// Busca os dados atuais para preservar campos opcionais.
	current, err := uc.ideogramRepo.FindByCharacter(ctx, character)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return nil, entity.ErrNotFound
		}
		return nil, fmt.Errorf("buscar ideograma para atualização: %w", err)
	}

	// Pinyin é obrigatório.
	ideog.Pinyin = strings.TrimSpace(ideog.Pinyin)
	if ideog.Pinyin == "" {
		return nil, entity.ErrEmptyPinyin
	}

	// Exige ao menos um significado válido.
	meanings := make([]string, 0, len(ideog.Meanings))
	for _, meaning := range ideog.Meanings {
		meaning = strings.TrimSpace(meaning)
		if meaning != "" {
			meanings = append(meanings, meaning)
		}
	}

	if len(meanings) == 0 {
		return nil, entity.ErrEmptyMeanings
	}
	ideog.Meanings = meanings

	// Campos opcionais: preservar os valores atuais quando omitidos.
	ideog.Radical = strings.TrimSpace(ideog.Radical)
	if ideog.Radical == "" {
		ideog.Radical = current.Radical
	}

	if ideog.StrokeCount == 0 {
		ideog.StrokeCount = current.StrokeCount
	}

	// O caractere original permanece como identificador.
	updated, err := uc.ideogramRepo.Update(ctx, character, ideog)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return nil, entity.ErrNotFound
		}
		return nil, fmt.Errorf("atualizar ideograma: %w", err)
	}

	return updated, nil
}
