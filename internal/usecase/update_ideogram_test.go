package usecase

import (
	"context"
	"errors"
	"testing"

	"mandarin-app/internal/entity"
)

type updateIdeogramRepoMock struct {
	current         *entity.Ideogram
	updated         *entity.Ideogram
	findErr         error
	updateErr       error
	findCalled      bool
	updateCalled    bool
	updateCharacter string
}

func (m *updateIdeogramRepoMock) FindByCharacter(
	ctx context.Context,
	char string,
) (*entity.Ideogram, error) {
	m.findCalled = true

	if m.findErr != nil {
		return nil, m.findErr
	}

	if m.current == nil {
		return nil, entity.ErrNotFound
	}

	return m.current, nil
}

func (m *updateIdeogramRepoMock) SearchByMeaning(
	ctx context.Context,
	meaning string,
) ([]entity.Ideogram, error) {
	return nil, nil
}

func (m *updateIdeogramRepoMock) Create(
	ctx context.Context,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	return nil, nil
}

func (m *updateIdeogramRepoMock) Update(
	ctx context.Context,
	character string,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	m.updateCalled = true
	m.updateCharacter = character

	if m.updateErr != nil {
		return nil, m.updateErr
	}

	if m.updated != nil {
		return m.updated, nil
	}

	result := *ideog
	result.Character = character
	result.ID = m.current.ID

	return &result, nil
}

func existingIdeogram() *entity.Ideogram {
	return &entity.Ideogram{
		ID:          1,
		Character:   "永",
		Pinyin:      "yǒng",
		Meanings:    []string{"eternidade"},
		Radical:     "水",
		StrokeCount: 5,
	}
}

func validUpdateIdeogram() *entity.Ideogram {
	return &entity.Ideogram{
		Pinyin:   "yǒng",
		Meanings: []string{" permanência ", " ", "duradouro"},
	}
}

func TestUpdateIdeogramUseCase_ExecuteSuccess(t *testing.T) {
	repo := &updateIdeogramRepoMock{
		current: existingIdeogram(),
	}
	uc := NewUpdateIdeogramUseCase(repo)

	result, err := uc.Execute(
		context.Background(),
		"永",
		validUpdateIdeogram(),
	)

	if err != nil {
		t.Fatalf("não esperava erro: %v", err)
	}

	if !repo.findCalled {
		t.Fatal("esperava que FindByCharacter fosse chamado")
	}

	if !repo.updateCalled {
		t.Fatal("esperava que Update fosse chamado")
	}

	if repo.updateCharacter != "永" {
		t.Errorf("identificador = %q; esperado %q",
			repo.updateCharacter, "永")
	}

	if result.Character != "永" {
		t.Errorf("Character = %q; esperado %q",
			result.Character, "永")
	}

	if result.Pinyin != "yǒng" {
		t.Errorf("Pinyin = %q; esperado %q",
			result.Pinyin, "yǒng")
	}

	if len(result.Meanings) != 2 {
		t.Fatalf("esperava 2 significados; recebeu %d",
			len(result.Meanings))
	}

	if result.Meanings[0] != "permanência" ||
		result.Meanings[1] != "duradouro" {
		t.Errorf("significados inesperados: %v", result.Meanings)
	}

	if result.Radical != "水" {
		t.Errorf("Radical = %q; esperado %q", result.Radical, "水")
	}

	if result.StrokeCount != 5 {
		t.Errorf("StrokeCount = %d; esperado %d",
			result.StrokeCount, 5)
	}
}

func TestUpdateIdeogramUseCase_ExecuteNotFound(t *testing.T) {
	repo := &updateIdeogramRepoMock{}
	uc := NewUpdateIdeogramUseCase(repo)

	_, err := uc.Execute(
		context.Background(),
		"永",
		validUpdateIdeogram(),
	)

	if !errors.Is(err, entity.ErrNotFound) {
		t.Errorf("esperava ErrNotFound; recebeu %v", err)
	}

	if repo.updateCalled {
		t.Error("Update não deveria ser chamado se o ideograma não existe")
	}
}

func TestUpdateIdeogramUseCase_ExecuteEmptyCharacter(t *testing.T) {
	repo := &updateIdeogramRepoMock{
		current: existingIdeogram(),
	}
	uc := NewUpdateIdeogramUseCase(repo)

	_, err := uc.Execute(
		context.Background(),
		"   ",
		validUpdateIdeogram(),
	)

	if !errors.Is(err, entity.ErrEmptyCharacter) {
		t.Errorf("esperava ErrEmptyCharacter; recebeu %v", err)
	}

	if repo.findCalled || repo.updateCalled {
		t.Error("repositório não deveria ser chamado")
	}
}

func TestUpdateIdeogramUseCase_ExecuteEmptyPinyin(t *testing.T) {
	repo := &updateIdeogramRepoMock{
		current: existingIdeogram(),
	}
	uc := NewUpdateIdeogramUseCase(repo)

	ideog := validUpdateIdeogram()
	ideog.Pinyin = "   "

	_, err := uc.Execute(context.Background(), "永", ideog)

	if !errors.Is(err, entity.ErrEmptyPinyin) {
		t.Errorf("esperava ErrEmptyPinyin; recebeu %v", err)
	}

	if repo.updateCalled {
		t.Error("Update não deveria ser chamado com pinyin vazio")
	}
}

func TestUpdateIdeogramUseCase_ExecuteEmptyMeanings(t *testing.T) {
	repo := &updateIdeogramRepoMock{
		current: existingIdeogram(),
	}
	uc := NewUpdateIdeogramUseCase(repo)

	ideog := validUpdateIdeogram()
	ideog.Meanings = []string{" ", "   "}

	_, err := uc.Execute(context.Background(), "永", ideog)

	if !errors.Is(err, entity.ErrEmptyMeanings) {
		t.Errorf("esperava ErrEmptyMeanings; recebeu %v", err)
	}

	if repo.updateCalled {
		t.Error("Update não deveria ser chamado sem significados válidos")
	}
}

func TestUpdateIdeogramUseCase_ExecutePreservesOptionalFields(t *testing.T) {
	repo := &updateIdeogramRepoMock{
		current: existingIdeogram(),
	}
	uc := NewUpdateIdeogramUseCase(repo)

	ideog := validUpdateIdeogram()

	_, err := uc.Execute(context.Background(), "永", ideog)
	if err != nil {
		t.Fatalf("não esperava erro: %v", err)
	}

	if ideog.Radical != "水" {
		t.Errorf("Radical = %q; esperado que fosse preservado como %q",
			ideog.Radical, "水")
	}

	if ideog.StrokeCount != 5 {
		t.Errorf("StrokeCount = %d; esperado que fosse preservado como 5",
			ideog.StrokeCount)
	}
}

func TestUpdateIdeogramUseCase_ExecuteRepositoryUpdateError(t *testing.T) {
	repoErr := errors.New("falha simulada no banco")
	repo := &updateIdeogramRepoMock{
		current:   existingIdeogram(),
		updateErr: repoErr,
	}
	uc := NewUpdateIdeogramUseCase(repo)

	_, err := uc.Execute(
		context.Background(),
		"永",
		validUpdateIdeogram(),
	)

	if !errors.Is(err, repoErr) {
		t.Errorf("esperava que o erro original fosse preservado; recebeu %v", err)
	}
}

func TestUpdateIdeogramUseCase_ExecuteNilIdeogram(t *testing.T) {
	repo := &updateIdeogramRepoMock{}
	uc := NewUpdateIdeogramUseCase(repo)

	_, err := uc.Execute(context.Background(), "永", nil)

	if err == nil {
		t.Fatal("esperava erro para ideograma nil")
	}

	if repo.findCalled || repo.updateCalled {
		t.Error("repositório não deveria ser chamado")
	}
}
