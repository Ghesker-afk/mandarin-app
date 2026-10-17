package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"mandarin-app/internal/entity"
	"mandarin-app/internal/usecase"
)

// Mock exclusivo dos testes do handler de cadastro.
type createIdeogramHandlerMockRepo struct {
	createErr error
	called    bool
}

func (m *createIdeogramHandlerMockRepo) FindByCharacter(
	ctx context.Context,
	char string,
) (*entity.Ideogram, error) {
	return nil, entity.ErrNotFound
}

func (m *createIdeogramHandlerMockRepo) SearchByMeaning(
	ctx context.Context,
	meaning string,
) ([]entity.Ideogram, error) {
	return nil, nil
}

func (m *createIdeogramHandlerMockRepo) Create(
	ctx context.Context,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	m.called = true

	if m.createErr != nil {
		return nil, m.createErr
	}

	ideog.ID = 1
	return ideog, nil
}

func (m *createIdeogramHandlerMockRepo) Update(
	ctx context.Context,
	character string,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	return nil, entity.ErrNotFound
}

func (m *createIdeogramHandlerMockRepo) FindAll(
	ctx context.Context,
) ([]entity.Ideogram, error) {
	return []entity.Ideogram{}, nil
}

func newCreateIdeogramTestHandler(
	repo *createIdeogramHandlerMockRepo,
) http.Handler {
	uc := usecase.NewCreateIdeogramUseCase(repo)
	return NewCreateIdeogramHandler(uc)
}

func TestCreateIdeogramHandler_Success(t *testing.T) {
	repo := &createIdeogramHandlerMockRepo{}
	handler := newCreateIdeogramTestHandler(repo)

	body := `{
		"character": "永",
		"pinyin": "yǒng",
		"meanings": ["eternidade", "permanência"],
		"radical": "水",
		"stroke_count": 5
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/ideograms",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"status = %d; esperado %d; resposta: %s",
			rec.Code,
			http.StatusCreated,
			rec.Body.String(),
		)
	}

	if contentType := rec.Header().Get("Content-Type"); !strings.HasPrefix(
		contentType,
		"application/json",
	) {
		t.Errorf("Content-Type inesperado: %q", contentType)
	}

	if !repo.called {
		t.Error("esperava que o repositório fosse chamado")
	}

	if !strings.Contains(rec.Body.String(), `"id":1`) {
		t.Errorf("resposta deveria conter o ID criado: %s", rec.Body.String())
	}
}

func TestCreateIdeogramHandler_InvalidJSON(t *testing.T) {
	repo := &createIdeogramHandlerMockRepo{}
	handler := newCreateIdeogramTestHandler(repo)

	req := httptest.NewRequest(
		http.MethodPost,
		"/ideograms",
		strings.NewReader(`{"character":`),
	)

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf(
			"status = %d; esperado %d",
			rec.Code,
			http.StatusBadRequest,
		)
	}

	if repo.called {
		t.Error("repositório não deveria ser chamado")
	}
}

func TestCreateIdeogramHandler_ValidationErrors(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{
			name: "caractere vazio",
			body: `{
				"character": " ",
				"pinyin": "yǒng",
				"meanings": ["eternidade"]
			}`,
		},
		{
			name: "pinyin vazio",
			body: `{
				"character": "永",
				"pinyin": " ",
				"meanings": ["eternidade"]
			}`,
		},
		{
			name: "significados vazios",
			body: `{
				"character": "永",
				"pinyin": "yǒng",
				"meanings": [" ", "  "]
			}`,
		},
		{
			name: "lista de significados vazia",
			body: `{
				"character": "永",
				"pinyin": "yǒng",
				"meanings": []
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createIdeogramHandlerMockRepo{}
			handler := newCreateIdeogramTestHandler(repo)

			req := httptest.NewRequest(
				http.MethodPost,
				"/ideograms",
				strings.NewReader(tt.body),
			)

			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Errorf(
					"status = %d; esperado %d; resposta: %s",
					rec.Code,
					http.StatusBadRequest,
					rec.Body.String(),
				)
			}

			if repo.called {
				t.Error("repositório não deveria ser chamado")
			}
		})
	}
}

func TestCreateIdeogramHandler_AlreadyExists(t *testing.T) {
	repo := &createIdeogramHandlerMockRepo{
		createErr: entity.ErrAlreadyExists,
	}
	handler := newCreateIdeogramTestHandler(repo)

	body := `{
		"character": "永",
		"pinyin": "yǒng",
		"meanings": ["eternidade"]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/ideograms",
		strings.NewReader(body),
	)

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusConflict {
		t.Errorf(
			"status = %d; esperado %d; resposta: %s",
			rec.Code,
			http.StatusConflict,
			rec.Body.String(),
		)
	}

	if !repo.called {
		t.Error("esperava que o repositório fosse chamado")
	}
}

func TestCreateIdeogramHandler_InternalError(t *testing.T) {
	repo := &createIdeogramHandlerMockRepo{
		createErr: errors.New("falha simulada no banco"),
	}
	handler := newCreateIdeogramTestHandler(repo)

	body := `{
		"character": "永",
		"pinyin": "yǒng",
		"meanings": ["eternidade"]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/ideograms",
		strings.NewReader(body),
	)

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf(
			"status = %d; esperado %d; resposta: %s",
			rec.Code,
			http.StatusInternalServerError,
			rec.Body.String(),
		)
	}
}
