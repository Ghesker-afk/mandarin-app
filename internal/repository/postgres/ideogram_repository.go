package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"mandarin-app/internal/entity"
)

type PostgresIdeogramRepository struct {
	db *pgx.Conn // Alterado de *pgxpool.Pool para *pgx.Conn
}

// NewPostgresIdeogramRepository recebe *pgx.Conn para combinar com o ConnectDB
func NewPostgresIdeogramRepository(db *pgx.Conn) *PostgresIdeogramRepository {
	return &PostgresIdeogramRepository{
		db: db,
	}
}

func (r *PostgresIdeogramRepository) FindByCharacter(ctx context.Context, char string) (*entity.Ideogram, error) {
	query := `
		SELECT id, character, pinyin, meanings, radical, stroke_count 
		FROM ideograms 
		WHERE character = $1
	`

	var ideog entity.Ideogram

	err := r.db.QueryRow(ctx, query, char).Scan(
		&ideog.ID,
		&ideog.Character,
		&ideog.Pinyin,
		&ideog.Meanings,
		&ideog.Radical,
		&ideog.StrokeCount,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, entity.ErrNotFound
		}
		return nil, fmt.Errorf("erro ao consultar ideograma no banco: %w", err)
	}

	return &ideog, nil
}

func (r *PostgresIdeogramRepository) SearchByMeaning(ctx context.Context, meaning string) ([]entity.Ideogram, error) {
	return []entity.Ideogram{}, nil
}
