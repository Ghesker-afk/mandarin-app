package usecase

import (
	"context"
	"errors"
	"testing"

	"mandarin-app/internal/entity"
)

type getIdeogramMockRepository struct {
	findFn func(context.Context, string) (*entity.Ideogram, error)
}

func (m *getIdeogramMockRepository) FindByCharacter(
	ctx context.Context,
	character string,
) (*entity.Ideogram, error) {
	return m.findFn(ctx, character)
}

func (m *getIdeogramMockRepository) SearchByMeaning(
	ctx context.Context,
	meaning string,
) ([]entity.Ideogram, error) {
	return nil, nil
}

func (m *getIdeogramMockRepository) Create(
	ctx context.Context,
	ideogram *entity.Ideogram,
) (*entity.Ideogram, error) {
	return nil, nil
}

func (m *getIdeogramMockRepository) Update(
	ctx context.Context,
	character string,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	return nil, entity.ErrNotFound
}

func TestGetIdeogramByCharacter(t *testing.T) {
	t.Run("retorna ideograma encontrado", func(t *testing.T) {
		expected := &entity.Ideogram{
			Character: "山",
			Pinyin:    "shān",
			Meanings:  []string{"montanha"},
		}

		repo := &getIdeogramMockRepository{
			findFn: func(
				ctx context.Context,
				character string,
			) (*entity.Ideogram, error) {
				if character != "山" {
					t.Fatalf("caractere recebido: %q", character)
				}

				return expected, nil
			},
		}

		uc := NewGetIdeogramByCharacterUseCase(repo)

		got, err := uc.Execute(context.Background(), " 山 ")

		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		if got.Character != expected.Character {
			t.Errorf(
				"caractere esperado %q, recebido %q",
				expected.Character,
				got.Character,
			)
		}
	})

	t.Run("retorna erro quando não encontra", func(t *testing.T) {
		repo := &getIdeogramMockRepository{
			findFn: func(
				context.Context,
				string,
			) (*entity.Ideogram, error) {
				return nil, entity.ErrNotFound
			},
		}

		uc := NewGetIdeogramByCharacterUseCase(repo)

		_, err := uc.Execute(context.Background(), "山")

		if !errors.Is(err, entity.ErrNotFound) {
			t.Fatalf("esperava ErrNotFound, recebeu: %v", err)
		}
	})

	t.Run("rejeita mais de um caractere", func(t *testing.T) {
		repo := &getIdeogramMockRepository{
			findFn: func(
				context.Context,
				string,
			) (*entity.Ideogram, error) {
				t.Fatal("o repositório não deveria ser chamado")
				return nil, nil
			},
		}

		uc := NewGetIdeogramByCharacterUseCase(repo)

		_, err := uc.Execute(context.Background(), "山水")

		if !errors.Is(err, entity.ErrInvalidCharacter) {
			t.Fatalf(
				"esperava ErrInvalidCharacter, recebeu: %v",
				err,
			)
		}
	})
}
