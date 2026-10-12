package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"mandarin-app/internal/entity"
	"mandarin-app/internal/usecase"
)

type CreateIdeogramHandler struct {
	useCase *usecase.CreateIdeogramUseCase
}

func NewCreateIdeogramHandler(
	uc *usecase.CreateIdeogramUseCase,
) *CreateIdeogramHandler {
	return &CreateIdeogramHandler{
		useCase: uc,
	}
}

func (h *CreateIdeogramHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	var ideog entity.Ideogram

	if err := json.NewDecoder(r.Body).Decode(&ideog); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	created, err := h.useCase.Execute(r.Context(), &ideog)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrEmptyCharacter):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, entity.ErrAlreadyExists):
			http.Error(w, "Ideograma já cadastrado", http.StatusConflict)
		case errors.Is(err, entity.ErrEmptyPinyin),
			errors.Is(err, entity.ErrEmptyMeanings):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			slog.Error("falha ao cadastrar ideograma", "error", err)
			http.Error(
				w,
				"Erro interno ao cadastrar ideograma",
				http.StatusInternalServerError,
			)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(created); err != nil {
		slog.Error("falha ao escrever resposta do cadastro", "error", err)
	}
}
