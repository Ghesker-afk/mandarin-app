package usecase_test

import (
	"context"
	"errors"
	"testing"

	// Substitua "mandarin-api" pelo nome exato que está no seu go.mod
	"mandarin-app/internal/entity"
	"mandarin-app/internal/usecase"
)

// 1. Criamos um Mock do Repositório que simula o banco de dados em memória
type MockIdeogramRepository struct {
	mockData    map[string]entity.Ideogram
	shouldError bool
}

func (m *MockIdeogramRepository) FindByCharacter(ctx context.Context, char string) (*entity.Ideogram, error) {
	if m.shouldError {
		return nil, errors.New("erro simulado de banco de dados")
	}

	ideogram, exists := m.mockData[char]
	if !exists {
		return nil, entity.ErrNotFound
	}

	return &ideogram, nil
}

func (m *MockIdeogramRepository) SearchByMeaning(ctx context.Context, meaning string) ([]entity.Ideogram, error) {
	return nil, nil
}

func (m *MockIdeogramRepository) Update(
	ctx context.Context,
	character string,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	return nil, entity.ErrNotFound
}

// 2. Teste Unitário da Regra de Negócio de Tradução
func TestTranslateIdeogramUseCase_Execute(t *testing.T) {
	// Preparamos dados falsos que estariam no "banco"
	fakeRepo := &MockIdeogramRepository{
		mockData: map[string]entity.Ideogram{
			"永": {
				ID:          1,
				Character:   "永",
				Pinyin:      "yǒng",
				Meanings:    []string{"eternidade", "permanência"},
				Radical:     "水",
				StrokeCount: 5,
			},
		},
	}

	// Instanciamos o UseCase passando o nosso repositório mockado
	uc := usecase.NewTranslateIdeogramUseCase(fakeRepo)

	// Executamos a tradução com um texto de teste (ex: "永")
	result, err := uc.Execute(context.Background(), "永")

	// Validações
	if err != nil {
		t.Fatalf("não esperava erro, mas obteve: %v", err)
	}

	if result.OriginalText != "永" {
		t.Errorf("esperava o texto original '永', obteve '%s'", result.OriginalText)
	}

	if len(result.Ideograms) != 1 {
		t.Errorf("esperava encontrar 1 ideograma, obteve %d", len(result.Ideograms))
	}

	if result.LiteralText != "eternidade" {
		t.Errorf("esperava o significado literal 'eternidade', obteve '%s'", result.LiteralText)
	}
}

// 3. Teste Unitário para validar cenários de erro (ex: string vazia)
func TestTranslateIdeogramUseCase_EmptyText(t *testing.T) {
	fakeRepo := &MockIdeogramRepository{}
	uc := usecase.NewTranslateIdeogramUseCase(fakeRepo)

	_, err := uc.Execute(context.Background(), "   ")

	if !errors.Is(err, entity.ErrEmptyCharacter) {
		t.Errorf("esperava o erro ErrEmptyCharacter, obteve: %v", err)
	}
}

func (m *MockIdeogramRepository) Create(
	ctx context.Context,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	return ideog, nil
}
