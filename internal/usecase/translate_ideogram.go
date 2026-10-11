package usecase

import (
	"context"
	"strings"

	"mandarin-app/internal/entity"
)

type TranslateIdeogramUseCase struct {
	repo entity.IdeogramRepository
}

func NewTranslateIdeogramUseCase(repo entity.IdeogramRepository) *TranslateIdeogramUseCase {
	return &TranslateIdeogramUseCase{
		repo: repo,
	}
}

func (uc *TranslateIdeogramUseCase) Execute(ctx context.Context, text string) (*entity.TranslationResult, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil, entity.ErrEmptyCharacter
	}

	var foundIdeograms []entity.Ideogram
	var literalMeanings []string

	// Itera sobre cada caractere/rune da string (suporte nativo a UTF-8/Mandarim no Go)
	for _, charRune := range text {
		charStr := string(charRune)

		ideogram, err := uc.repo.FindByCharacter(ctx, charStr)
		if err != nil {
			// Se um ideograma não for encontrado, podemos continuar ou sinalizar
			continue
		}

		foundIdeograms = append(foundIdeograms, *ideogram)
		if len(ideogram.Meanings) > 0 {
			literalMeanings = append(literalMeanings, ideogram.Meanings[0])
		}
	}

	result := &entity.TranslationResult{
		OriginalText: text,
		Ideograms:    foundIdeograms,
		LiteralText:  strings.Join(literalMeanings, " "),
	}

	return result, nil
}
