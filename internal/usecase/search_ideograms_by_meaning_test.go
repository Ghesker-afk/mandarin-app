package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"mandarin-app/internal/entity"
)

type searchMeaningMockRepository struct {
	searchFn func(context.Context, string) ([]entity.Ideogram, error)
}

func (m *searchMeaningMockRepository) FindByCharacter(
	ctx context.Context,
	character string,
) (*entity.Ideogram, error) {
	return nil, nil
}

func (m *searchMeaningMockRepository) SearchByMeaning(
	ctx context.Context,
	meaning string,
) ([]entity.Ideogram, error) {
	return m.searchFn(ctx, meaning)
}

func (m *searchMeaningMockRepository) Create(
	ctx context.Context,
	ideogram *entity.Ideogram,
) (*entity.Ideogram, error) {
	return nil, nil
}

func (m *searchMeaningMockRepository) Update(
	ctx context.Context,
	character string,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	return nil, entity.ErrNotFound
}

func (m *searchMeaningMockRepository) FindAll(
	ctx context.Context,
) ([]entity.Ideogram, error) {
	return []entity.Ideogram{}, nil
}

func TestSearchIdeogramsByMeaning(t *testing.T) {
	t.Run("retorna ideogramas encontrados", func(t *testing.T) {
		expected := []entity.Ideogram{
			{
				Character: "山",
				Pinyin:    "shān",
				Meanings:  []string{"montanha"},
			},
		}

		repo := &searchMeaningMockRepository{
			searchFn: func(
				ctx context.Context,
				meaning string,
			) ([]entity.Ideogram, error) {
				if meaning != "montanha" {
					t.Errorf("significado recebido: %q", meaning)
				}

				return expected, nil
			},
		}

		uc := NewSearchIdeogramsByMeaningUseCase(repo)

		got, err := uc.Execute(context.Background(), " montanha ")

		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("resultado inesperado: got %v, want %v", got, expected)
		}
	})

	t.Run("rejeita significado vazio", func(t *testing.T) {
		repo := &searchMeaningMockRepository{
			searchFn: func(
				context.Context,
				string,
			) ([]entity.Ideogram, error) {
				t.Fatal("o repositório não deveria ser chamado")
				return nil, nil
			},
		}

		uc := NewSearchIdeogramsByMeaningUseCase(repo)

		_, err := uc.Execute(context.Background(), "   ")

		if !errors.Is(err, entity.ErrEmptyMeaning) {
			t.Fatalf("esperava ErrEmptyMeaning, recebeu: %v", err)
		}
	})

	t.Run("propaga erro do repositório", func(t *testing.T) {
		repoErr := errors.New("falha simulada no banco")

		repo := &searchMeaningMockRepository{
			searchFn: func(
				context.Context,
				string,
			) ([]entity.Ideogram, error) {
				return nil, repoErr
			},
		}

		uc := NewSearchIdeogramsByMeaningUseCase(repo)

		_, err := uc.Execute(context.Background(), "montanha")

		if !errors.Is(err, repoErr) {
			t.Fatalf("esperava erro do repositório, recebeu: %v", err)
		}
	})
}
