package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"mandarin-app/internal/entity"
	"mandarin-app/internal/usecase"
)

type UpdateIdeogramHandler struct {
	updateIdeogramUseCase *usecase.UpdateIdeogramUseCase
}

func NewUpdateIdeogramHandler(
	uc *usecase.UpdateIdeogramUseCase,
) *UpdateIdeogramHandler {
	return &UpdateIdeogramHandler{
		updateIdeogramUseCase: uc,
	}
}

func (h *UpdateIdeogramHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	character := r.PathValue("character")

	var ideog entity.Ideogram

	if err := json.NewDecoder(r.Body).Decode(&ideog); err != nil {
		http.Error(w, `{"error":"JSON inválido"}`, http.StatusBadRequest)
		return
	}

	updated, err := h.updateIdeogramUseCase.Execute(
		r.Context(),
		character,
		&ideog,
	)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			http.Error(w, `{"error":"Ideograma não encontrado"}`, http.StatusNotFound)
		case errors.Is(err, entity.ErrEmptyCharacter),
			errors.Is(err, entity.ErrEmptyPinyin),
			errors.Is(err, entity.ErrEmptyMeanings):
			http.Error(w, `{"error":"Dados inválidos"}`, http.StatusBadRequest)
		default:
			http.Error(w, `{"error":"Erro interno ao atualizar ideograma"}`, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(updated); err != nil {
		// A resposta já foi iniciada; não é possível trocar o status HTTP.
		return
	}
}
