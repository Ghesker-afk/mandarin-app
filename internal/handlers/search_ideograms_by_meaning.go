package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"mandarin-app/internal/entity"
	"mandarin-app/internal/usecase"
)

type SearchIdeogramsByMeaningHandler struct {
	searchUseCase *usecase.SearchIdeogramsByMeaningUseCase
}

func NewSearchIdeogramsByMeaningHandler(
	uc *usecase.SearchIdeogramsByMeaningUseCase,
) *SearchIdeogramsByMeaningHandler {
	return &SearchIdeogramsByMeaningHandler{
		searchUseCase: uc,
	}
}

func (h *SearchIdeogramsByMeaningHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	meaning := r.URL.Query().Get("meaning")

	ideograms, err := h.searchUseCase.Execute(
		r.Context(),
		meaning,
	)
	if err != nil {
		if errors.Is(err, entity.ErrEmptyMeaning) {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		slog.Error(
			"falha ao buscar ideogramas por significado",
			"meaning", meaning,
			"error", err,
		)

		http.Error(
			w,
			"Erro interno ao buscar ideogramas",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(ideograms); err != nil {
		slog.Error(
			"falha ao serializar resultado da busca",
			"error", err,
		)
	}
}
