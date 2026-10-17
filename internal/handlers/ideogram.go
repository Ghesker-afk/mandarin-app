package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"mandarin-app/internal/entity"
	"mandarin-app/internal/usecase"
)

type IdeogramHandler struct {
	translateUseCase     *usecase.TranslateIdeogramUseCase
	listIdeogramsUseCase *usecase.ListIdeogramsUseCase
}

// NewIdeogramHandler injeta o caso de uso no handler
func NewIdeogramHandler(translateUseCase *usecase.TranslateIdeogramUseCase, listIdeogramsUseCase *usecase.ListIdeogramsUseCase) *IdeogramHandler {
	return &IdeogramHandler{
		translateUseCase:     translateUseCase,
		listIdeogramsUseCase: listIdeogramsUseCase,
	}
}

// TranslateHandler lida com a rota GET /translate?text=... ou POST de tradução
func (h *IdeogramHandler) TranslateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Extrai o texto da query string (ex: /translate?text=永)
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "O parâmetro 'text' é obrigatório", http.StatusBadRequest)
		return
	}

	// Executa a regra de negócio através do UseCase
	result, err := h.translateUseCase.Execute(r.Context(), text)
	if err != nil {
		if errors.Is(err, entity.ErrEmptyCharacter) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		slog.Error("falha ao processar tradução", "error", err)
		http.Error(
			w,
			"Erro interno ao processar tradução",
			http.StatusInternalServerError,
		)
		return
	}

	// Retorna o resultado estruturado em JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (h *IdeogramHandler) ListIdeogramsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método não permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	ideograms, err := h.listIdeogramsUseCase.Execute(r.Context())
	if err != nil {
		slog.Error("falha ao listar ideogramas", "error", err)
		http.Error(
			w,
			"Erro interno ao listar ideogramas",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(ideograms); err != nil {
		slog.Error("falha ao serializar ideogramas", "error", err)
	}
}
