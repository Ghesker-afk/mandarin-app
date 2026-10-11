package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type Ideogram struct {
	ID              int    `json:"id"`
	Character       string `json:"character"`
	Pinyin          string `json:"pinyin"`
	Meaning         string `json:"meaning"`
	ExampleSentence string `json:"example_sentence"`
}

type IdeogramHandler struct {
	DB *pgx.Conn
}

func NewIdeogramHandler(db *pgx.Conn) *IdeogramHandler {
	return &IdeogramHandler{DB: db}
}

// Handler para criar um novo ideograma (POST)
func (h *IdeogramHandler) CreateIdeogram(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var item Ideogram
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO ideograms (character, pinyin, meaning, example_sentence) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err = h.DB.QueryRow(context.Background(), query, item.Character, item.Pinyin, item.Meaning, item.ExampleSentence).Scan(&id)
	if err != nil {
		http.Error(w, "Erro ao salvar no banco de dados", http.StatusInternalServerError)
		return
	}

	item.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// Handler para listar todos os ideogramas (GET)
func (h *IdeogramHandler) ListIdeograms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	rows, err := h.DB.Query(context.Background(), `SELECT id, character, pinyin, meaning, example_sentence FROM ideograms ORDER BY id DESC`)
	if err != nil {
		http.Error(w, "Erro ao buscar ideogramas", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ideograms []Ideogram
	for rows.Next() {
		var i Ideogram
		if err := rows.Scan(&i.ID, &i.Character, &i.Pinyin, &i.Meaning, &i.ExampleSentence); err != nil {
			continue
		}
		ideograms = append(ideograms, i)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ideograms)
}
