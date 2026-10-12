package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"mandarin-app/internal/entity"
	"mandarin-app/internal/usecase"
)

type GetIdeogramByCharacterHandler struct {
	getIdeogramUseCase *usecase.GetIdeogramByCharacterUseCase
}

func NewGetIdeogramByCharacterHandler(
	uc *usecase.GetIdeogramByCharacterUseCase,
) *GetIdeogramByCharacterHandler {
	return &GetIdeogramByCharacterHandler{
		getIdeogramUseCase: uc,
	}
}

func (h *GetIdeogramByCharacterHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	character := r.PathValue("character")

	ideogram, err := h.getIdeogramUseCase.Execute(
		r.Context(),
		character,
	)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrEmptyCharacter),
			errors.Is(err, entity.ErrInvalidCharacter):
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		case errors.Is(err, entity.ErrNotFound):
			http.Error(w, "Ideograma não encontrado", http.StatusNotFound)
			return

		default:
			slog.Error(
				"falha ao consultar ideograma",
				"character", character,
				"error", err,
			)
			http.Error(
				w,
				"Erro interno ao consultar ideograma",
				http.StatusInternalServerError,
			)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(ideogram); err != nil {
		slog.Error("falha ao serializar ideograma", "error", err)
	}
}
