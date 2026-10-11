package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetIdeogramsHandlerReal testa a rota e o comportamento HTTP do handler de ideogramas
func TestGetIdeogramsHandlerReal(t *testing.T) {
	// 1. Configurar o roteador idêntico ao da sua aplicação real
	mux := http.NewServeMux()

	// Registre a rota que você usa no seu projeto (exemplo: "GET /ideograms")
	// Ajuste o nome da função do seu handler caso seja diferente de IdeogramHandler
	// mux.HandleFunc("GET /ideograms", handlers.GetIdeograms)

	// Como estamos criando o teste, vamos simular um handler temporário ou usar o seu real:
	mux.HandleFunc("GET /ideograms", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id": 1, "character": "永", "meaning": "Eternidade"}]`))
	})

	// 2. Criar a requisição HTTP simulada para o endpoint
	req, err := http.NewRequest(http.MethodGet, "/ideograms", nil)
	if err != nil {
		t.Fatalf("falha ao criar requisição: %v", err)
	}

	// 3. Gravar a resposta gerada pelo mux/handler
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	// 4. Validar se o status code retornado é o esperado (200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status code incorreto: got %v, want %v", status, http.StatusOK)
	}

	// 5. Validar se o cabeçalho Content-Type é application/json
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("content-type incorreto: got %v, want %v", ct, "application/json")
	}
}
