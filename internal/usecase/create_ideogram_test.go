package usecase

import (
	"context"
	"errors"
	"testing"

	"mandarin-app/internal/entity"
)

type createIdeogramRepoMock struct {
	createErr error
	created   *entity.Ideogram
	called    bool
}

func (m *createIdeogramRepoMock) FindByCharacter(
	ctx context.Context,
	char string,
) (*entity.Ideogram, error) {
	return nil, entity.ErrNotFound
}

func (m *createIdeogramRepoMock) SearchByMeaning(
	ctx context.Context,
	meaning string,
) ([]entity.Ideogram, error) {
	return nil, nil
}

func (m *createIdeogramRepoMock) Create(
	ctx context.Context,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	m.called = true

	if m.createErr != nil {
		return nil, m.createErr
	}

	if m.created != nil {
		return m.created, nil
	}

	return ideog, nil
}

func validIdeogram() *entity.Ideogram {
	return &entity.Ideogram{
		Character:   " 永 ",
		Pinyin:      " yǒng ",
		Meanings:    []string{" eternidade ", " ", "permanência"},
		Radical:     " 水 ",
		StrokeCount: 5,
	}
}

func TestCreateIdeogramUseCase_ExecuteSuccess(t *testing.T) {
	repo := &createIdeogramRepoMock{}
	uc := NewCreateIdeogramUseCase(repo)

	result, err := uc.Execute(context.Background(), validIdeogram())

	if err != nil {
		t.Fatalf("não esperava erro: %v", err)
	}

	if !repo.called {
		t.Fatal("esperava que o repositório fosse chamado")
	}

	if result.Character != "永" {
		t.Errorf("Character = %q; esperado %q", result.Character, "永")
	}

	if result.Pinyin != "yǒng" {
		t.Errorf("Pinyin = %q; esperado %q", result.Pinyin, "yǒng")
	}

	if result.Radical != "水" {
		t.Errorf("Radical = %q; esperado %q", result.Radical, "水")
	}

	if len(result.Meanings) != 2 {
		t.Fatalf("esperava 2 significados; recebeu %d", len(result.Meanings))
	}

	if result.Meanings[0] != "eternidade" ||
		result.Meanings[1] != "permanência" {
		t.Errorf("significados inesperados: %v", result.Meanings)
	}
}

func TestCreateIdeogramUseCase_ExecuteEmptyCharacter(t *testing.T) {
	repo := &createIdeogramRepoMock{}
	uc := NewCreateIdeogramUseCase(repo)

	ideog := validIdeogram()
	ideog.Character = "   "

	_, err := uc.Execute(context.Background(), ideog)

	if !errors.Is(err, entity.ErrEmptyCharacter) {
		t.Errorf("esperava ErrEmptyCharacter; recebeu %v", err)
	}

	if repo.called {
		t.Error("repositório não deveria ser chamado")
	}
}

func TestCreateIdeogramUseCase_ExecuteEmptyPinyin(t *testing.T) {
	repo := &createIdeogramRepoMock{}
	uc := NewCreateIdeogramUseCase(repo)

	ideog := validIdeogram()
	ideog.Pinyin = "   "

	_, err := uc.Execute(context.Background(), ideog)

	if err == nil {
		t.Fatal("esperava erro para pinyin vazio")
	}

	if repo.called {
		t.Error("repositório não deveria ser chamado")
	}
}

func TestCreateIdeogramUseCase_ExecuteEmptyMeanings(t *testing.T) {
	repo := &createIdeogramRepoMock{}
	uc := NewCreateIdeogramUseCase(repo)

	ideog := validIdeogram()
	ideog.Meanings = []string{" ", "   "}

	_, err := uc.Execute(context.Background(), ideog)

	if err == nil {
		t.Fatal("esperava erro quando todos os significados estão vazios")
	}

	if repo.called {
		t.Error("repositório não deveria ser chamado")
	}
}

func TestCreateIdeogramUseCase_ExecuteAlreadyExists(t *testing.T) {
	repo := &createIdeogramRepoMock{
		createErr: entity.ErrAlreadyExists,
	}
	uc := NewCreateIdeogramUseCase(repo)

	_, err := uc.Execute(context.Background(), validIdeogram())

	if !errors.Is(err, entity.ErrAlreadyExists) {
		t.Errorf("esperava ErrAlreadyExists; recebeu %v", err)
	}
}

func TestCreateIdeogramUseCase_ExecuteRepositoryError(t *testing.T) {
	repoErr := errors.New("falha simulada no banco")

	repo := &createIdeogramRepoMock{
		createErr: repoErr,
	}
	uc := NewCreateIdeogramUseCase(repo)

	_, err := uc.Execute(context.Background(), validIdeogram())

	if !errors.Is(err, repoErr) {
		t.Errorf("esperava que o erro original fosse preservado; recebeu %v", err)
	}
}

func TestCreateIdeogramUseCase_ExecuteNilIdeogram(t *testing.T) {
	repo := &createIdeogramRepoMock{}
	uc := NewCreateIdeogramUseCase(repo)

	_, err := uc.Execute(context.Background(), nil)

	if err == nil {
		t.Fatal("esperava erro para ideograma nil")
	}

	if repo.called {
		t.Error("repositório não deveria ser chamado")
	}
}
