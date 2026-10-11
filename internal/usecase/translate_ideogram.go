
package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"mandarin-app/internal/entity"
)

type TranslateIdeogramUseCase struct {
	repo entity.IdeogramRepository
}

func NewTranslateIdeogramUseCase(
	repo entity.IdeogramRepository,
) *TranslateIdeogramUseCase {
	return &TranslateIdeogramUseCase{
		repo: repo,
	}
}

func (uc *TranslateIdeogramUseCase) Execute(
	ctx context.Context,
	text string,
) (*entity.TranslationResult, error) {
	text = strings.TrimSpace(text)

	if text == "" {
		return nil, entity.ErrEmptyCharacter
	}

	foundIdeograms := make([]entity.Ideogram, 0)
	literalMeanings := make([]string, 0)

	for _, charRune := range text {
		char := string(charRune)

		ideogram, err := uc.repo.FindByCharacter(ctx, char)

		if err != nil {
			if errors.Is(err, entity.ErrNotFound) {
				continue
			}

			return nil, fmt.Errorf(
				"buscar ideograma %q: %w",
				char,
				err,
			)
		}

		foundIdeograms = append(foundIdeograms, *ideogram)

		if len(ideogram.Meanings) > 0 {
			literalMeanings = append(
				literalMeanings,
				ideogram.Meanings[0],
			)
		}
	}

	return &entity.TranslationResult{
		OriginalText: text,
		Ideograms:    foundIdeograms,
		LiteralText:  strings.Join(literalMeanings, " "),
	}, nil
}
