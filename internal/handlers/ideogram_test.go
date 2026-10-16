package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	// Substitua "mandarin-api" pelo nome exato do seu go.mod
	"mandarin-app/internal/entity"
	"mandarin-app/internal/handlers"
	"mandarin-app/internal/usecase"
)

// Mock do repositório apenas para testar o handler
type mockRepoForHandler struct{}

func (m *mockRepoForHandler) FindByCharacter(ctx context.Context, char string) (*entity.Ideogram, error) {
	return &entity.Ideogram{
		ID:          1,
		Character:   "永",
		Pinyin:      "yǒng",
		Meanings:    []string{"eternidade"},
		Radical:     "水",
		StrokeCount: 5,
	}, nil
}

func (m *mockRepoForHandler) SearchByMeaning(ctx context.Context, meaning string) ([]entity.Ideogram, error) {
	return nil, nil
}

// TestTranslateHandlerReal testa a rota e o comportamento HTTP do handler refatorado
func TestTranslateHandlerReal(t *testing.T) {
	// 1. Montamos a cadeia de dependências (Repositório Falso -> UseCase -> Handler)
	repo := &mockRepoForHandler{}
	uc := usecase.NewTranslateIdeogramUseCase(repo)
	handler := handlers.NewIdeogramHandler(uc)

	// 2. Configurar o roteador idêntico ao da aplicação real
	mux := http.NewServeMux()
	mux.HandleFunc("GET /translate", handler.TranslateHandler)

	// 3. Criar a requisição HTTP simulada passando o parâmetro text
	req, err := http.NewRequest(http.MethodGet, "/translate?text=永", nil)
	if err != nil {
		t.Fatalf("falha ao criar requisição: %v", err)
	}

	// 4. Gravar a resposta gerada pelo mux/handler
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	// 5. Validar se o status code retornado é o esperado (200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status code incorreto: got %v, want %v", status, http.StatusOK)
	}

	// 6. Validar se o cabeçalho Content-Type é application/json
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("content-type incorreto: got %v, want %v", ct, "application/json")
	}
}

func (m *mockRepoForHandler) Create(
	ctx context.Context,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	return ideog, nil
}

func (m *mockRepoForHandler) Update(
	ctx context.Context,
	character string,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	return nil, entity.ErrNotFound
}
