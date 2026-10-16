package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"mandarin-app/internal/entity"
	"mandarin-app/internal/usecase"
)

type updateHandlerRepoMock struct {
	current          *entity.Ideogram
	updated          *entity.Ideogram
	updatedCharacter string
	updateCalled     bool
}

func (m *updateHandlerRepoMock) FindByCharacter(
	ctx context.Context,
	char string,
) (*entity.Ideogram, error) {
	if m.current == nil {
		return nil, entity.ErrNotFound
	}

	return m.current, nil
}

func (m *updateHandlerRepoMock) SearchByMeaning(
	ctx context.Context,
	meaning string,
) ([]entity.Ideogram, error) {
	return nil, nil
}

func (m *updateHandlerRepoMock) Create(
	ctx context.Context,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	return nil, nil
}

func (m *updateHandlerRepoMock) Update(
	ctx context.Context,
	character string,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	m.updateCalled = true
	m.updatedCharacter = character

	if m.updated != nil {
		return m.updated, nil
	}

	result := *ideog
	result.ID = m.current.ID
	result.Character = character

	return &result, nil
}

func TestUpdateIdeogramHandler_Success(t *testing.T) {
	// Arrange: preparar as dependências e os dados.
	repo := &updateHandlerRepoMock{
		current: &entity.Ideogram{
			ID:          1,
			Character:   "山",
			Pinyin:      "shan",
			Meanings:    []string{"montanha"},
			Radical:     "山",
			StrokeCount: 3,
		},
	}

	updateUseCase := usecase.NewUpdateIdeogramUseCase(repo)
	handler := NewUpdateIdeogramHandler(updateUseCase)

	mux := http.NewServeMux()
	mux.Handle("PUT /ideograms/{character}", handler)

	body := map[string]interface{}{
		"pinyin":       "shān",
		"meanings":     []string{"montanha", "monte"},
		"radical":      "山",
		"stroke_count": 3,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("erro ao criar JSON de teste: %v", err)
	}

	req := httptest.NewRequest(
		http.MethodPut,
		"/ideograms/山",
		bytes.NewReader(jsonBody),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	// Act: executar a requisição.
	mux.ServeHTTP(rec, req)

	// Assert: verificar o resultado.
	if rec.Code != http.StatusOK {
		t.Fatalf(
			"status esperado %d, recebido %d; resposta: %s",
			http.StatusOK,
			rec.Code,
			rec.Body.String(),
		)
	}

	if !repo.updateCalled {
		t.Fatal("esperava que o repositório fosse chamado para atualizar")
	}

	if repo.updatedCharacter != "山" {
		t.Errorf(
			"caractere esperado %q, recebido %q",
			"山",
			repo.updatedCharacter,
		)
	}

	var response entity.Ideogram
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("erro ao decodificar resposta JSON: %v", err)
	}

	if response.ID != 1 {
		t.Errorf("ID esperado 1, recebido %d", response.ID)
	}

	if response.Character != "山" {
		t.Errorf(
			"caractere esperado %q, recebido %q",
			"山",
			response.Character,
		)
	}

	if response.Pinyin != "shān" {
		t.Errorf(
			"pinyin esperado %q, recebido %q",
			"shān",
			response.Pinyin,
		)
	}

	if len(response.Meanings) != 2 {
		t.Fatalf(
			"esperava 2 significados, recebeu %d",
			len(response.Meanings),
		)
	}

	if response.Meanings[1] != "monte" {
		t.Errorf(
			"segundo significado esperado %q, recebido %q",
			"monte",
			response.Meanings[1],
		)
	}
}
