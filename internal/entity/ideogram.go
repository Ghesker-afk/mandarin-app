package entity

import "errors"

var (
	ErrEmptyCharacter = errors.New("o caractere ideográfico não pode ser vazio")
	ErrNotFound       = errors.New("ideograma não encontrado")
	ErrAlreadyExists  = errors.New("ideogram already exists")
	ErrEmptyPinyin    = errors.New("pinyin is required")
	ErrEmptyMeanings  = errors.New("at least one meaning is required")
)

// Ideogram representa a entidade de um ideograma mandarim no nosso sistema
type Ideogram struct {
	ID          int64    `json:"id"`
	Character   string   `json:"character"` // ex: "永"
	Pinyin      string   `json:"pinyin"`    // ex: "yǒng"
	Meanings    []string `json:"meanings"`  // ex: ["eternidade", "permanência"]
	Radical     string   `json:"radical"`   // ex: "水"
	StrokeCount int      `json:"stroke_count"`
}

// TranslationResult representa o resultado retornado pelo nosso motor de tradução
type TranslationResult struct {
	OriginalText string     `json:"original_text"`
	Ideograms    []Ideogram `json:"ideograms"`
	LiteralText  string     `json:"literal_text"`
}
