package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"mandarin-app/internal/entity"
	"mandarin-app/internal/usecase"
)

type searchMeaningHandlerMockRepo struct {
	searchFn func(context.Context, string) ([]entity.Ideogram, error)
}

func (m *searchMeaningHandlerMockRepo) FindByCharacter(
	context.Context,
	string,
) (*entity.Ideogram, error) {
	return nil, nil
}

func (m *searchMeaningHandlerMockRepo) SearchByMeaning(
	ctx context.Context,
	meaning string,
) ([]entity.Ideogram, error) {
	return m.searchFn(ctx, meaning)
}

func (m *searchMeaningHandlerMockRepo) Create(
	context.Context,
	*entity.Ideogram,
) (*entity.Ideogram, error) {
	return nil, nil
}

func (m *searchMeaningHandlerMockRepo) Update(
	ctx context.Context,
	character string,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	return nil, entity.ErrNotFound
}

func (m *searchMeaningHandlerMockRepo) FindAll(
	ctx context.Context,
) ([]entity.Ideogram, error) {
	return []entity.Ideogram{}, nil
}

func newSearchMeaningTestHandler(
	repo *searchMeaningHandlerMockRepo,
) http.Handler {
	uc := usecase.NewSearchIdeogramsByMeaningUseCase(repo)
	return NewSearchIdeogramsByMeaningHandler(uc)
}

func TestSearchIdeogramsByMeaningHandler(t *testing.T) {
	t.Run("retorna ideogramas encontrados", func(t *testing.T) {
		repo := &searchMeaningHandlerMockRepo{
			searchFn: func(
				_ context.Context,
				meaning string,
			) ([]entity.Ideogram, error) {
				if meaning != "montanha" {
					t.Errorf("significado recebido: %q", meaning)
				}

				return []entity.Ideogram{
					{
						Character:   "山",
						Pinyin:      "shān",
						Meanings:    []string{"montanha"},
						Radical:     "山",
						StrokeCount: 3,
					},
				}, nil
			},
		}

		handler := newSearchMeaningTestHandler(repo)
		req := httptest.NewRequest(
			http.MethodGet,
			"/ideograms/search?meaning=montanha",
			nil,
		)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("esperava status 200, recebeu %d", rec.Code)
		}

		if got := rec.Header().Get("Content-Type"); got != "application/json" {
			t.Errorf("Content-Type inesperado: %q", got)
		}

		var result []entity.Ideogram
		if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
			t.Fatalf("resposta não é JSON válido: %v", err)
		}

		if len(result) != 1 || result[0].Character != "山" {
			t.Errorf("resultado inesperado: %+v", result)
		}
	})

	t.Run("retorna lista vazia quando não encontra resultados", func(t *testing.T) {
		repo := &searchMeaningHandlerMockRepo{
			searchFn: func(
				context.Context,
				string,
			) ([]entity.Ideogram, error) {
				return []entity.Ideogram{}, nil
			},
		}

		handler := newSearchMeaningTestHandler(repo)
		req := httptest.NewRequest(
			http.MethodGet,
			"/ideograms/search?meaning=xyz",
			nil,
		)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("esperava status 200, recebeu %d", rec.Code)
		}

		if body := rec.Body.String(); body != "[]\n" {
			t.Errorf("esperava lista vazia, recebeu %q", body)
		}
	})

	t.Run("rejeita significado vazio", func(t *testing.T) {
		repo := &searchMeaningHandlerMockRepo{
			searchFn: func(
				context.Context,
				string,
			) ([]entity.Ideogram, error) {
				t.Fatal("repositório não deveria ser chamado")
				return nil, nil
			},
		}

		handler := newSearchMeaningTestHandler(repo)
		req := httptest.NewRequest(
			http.MethodGet,
			"/ideograms/search?meaning=%20%20",
			nil,
		)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("esperava status 400, recebeu %d", rec.Code)
		}
	})

	t.Run("retorna erro interno quando a busca falha", func(t *testing.T) {
		repo := &searchMeaningHandlerMockRepo{
			searchFn: func(
				context.Context,
				string,
			) ([]entity.Ideogram, error) {
				return nil, errors.New("falha simulada no banco")
			},
		}

		handler := newSearchMeaningTestHandler(repo)
		req := httptest.NewRequest(
			http.MethodGet,
			"/ideograms/search?meaning=montanha",
			nil,
		)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("esperava status 500, recebeu %d", rec.Code)
		}
	})
}
