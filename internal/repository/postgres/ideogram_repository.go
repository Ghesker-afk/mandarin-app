package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"mandarin-app/internal/entity"
)

type PostgresIdeogramRepository struct {
	db *pgx.Conn
}

func NewPostgresIdeogramRepository(db *pgx.Conn) *PostgresIdeogramRepository {
	return &PostgresIdeogramRepository{
		db: db,
	}
}

func (r *PostgresIdeogramRepository) FindAll(
	ctx context.Context,
) ([]entity.Ideogram, error) {
	query := `
        SELECT id, character, pinyin, COALESCE(meanings, '{}'), COALESCE(radical, ''), COALESCE(stroke_count, 0)
        FROM ideograms
        ORDER BY character;
    `

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf(
			"erro ao listar ideogramas: %w",
			err,
		)
	}
	defer rows.Close()

	ideograms := make([]entity.Ideogram, 0)

	for rows.Next() {
		var ideog entity.Ideogram

		err := rows.Scan(
			&ideog.ID,
			&ideog.Character,
			&ideog.Pinyin,
			&ideog.Meanings,
			&ideog.Radical,
			&ideog.StrokeCount,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"erro ao ler ideograma listado: %w",
				err,
			)
		}

		ideograms = append(ideograms, ideog)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"erro ao percorrer ideogramas: %w",
			err,
		)
	}

	return ideograms, nil
}

func (r *PostgresIdeogramRepository) FindByCharacter(ctx context.Context, char string) (*entity.Ideogram, error) {
	query := `
		SELECT id, character, pinyin, COALESCE(meanings, '{}'), COALESCE(radical, ''), COALESCE(stroke_count, 0) 
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

func (r *PostgresIdeogramRepository) SearchByMeaning(
	ctx context.Context,
	meaning string,
) ([]entity.Ideogram, error) {
	query := `
		SELECT id, character, pinyin, COALESCE(meanings, '{}'), COALESCE(radical, ''), COALESCE(stroke_count, 0)
		FROM ideograms
		WHERE EXISTS (
			SELECT 1
			FROM unnest(meanings) AS item
			WHERE item ILIKE '%' || $1 || '%'
		)
		ORDER BY character;
	`

	rows, err := r.db.Query(ctx, query, meaning)
	if err != nil {
		return nil, fmt.Errorf(
			"erro ao buscar ideogramas por significado: %w",
			err,
		)
	}
	defer rows.Close()

	ideograms := make([]entity.Ideogram, 0)

	for rows.Next() {
		var ideog entity.Ideogram

		err := rows.Scan(
			&ideog.ID,
			&ideog.Character,
			&ideog.Pinyin,
			&ideog.Meanings,
			&ideog.Radical,
			&ideog.StrokeCount,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"erro ao ler ideograma encontrado: %w",
				err,
			)
		}

		ideograms = append(ideograms, ideog)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"erro ao percorrer resultados da busca: %w",
			err,
		)
	}

	return ideograms, nil
}

func (r *PostgresIdeogramRepository) Create(ctx context.Context, ideog *entity.Ideogram) (*entity.Ideogram, error) {
	query := `
	INSERT INTO ideograms (
		character,
		pinyin,
		meanings,
		radical,
		stroke_count
	)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, character, pinyin, meanings, radical, stroke_count;
	`

	err := r.db.QueryRow(
		ctx,
		query,
		ideog.Character,
		ideog.Pinyin,
		ideog.Meanings,
		ideog.Radical,
		ideog.StrokeCount,
	).Scan(
		&ideog.ID,
		&ideog.Character,
		&ideog.Pinyin,
		&ideog.Meanings,
		&ideog.Radical,
		&ideog.StrokeCount,
	)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, entity.ErrAlreadyExists
		}

		return nil, fmt.Errorf("erro ao criar ideograma no banco: %w", err)
	}

	return ideog, nil
}

func (r *PostgresIdeogramRepository) Update(
	ctx context.Context,
	character string,
	ideog *entity.Ideogram,
) (*entity.Ideogram, error) {
	query := `
		UPDATE ideograms
		SET
			pinyin = $1,
			meanings = $2,
			radical = $3,
			stroke_count = $4
		WHERE character = $5
		RETURNING id, character, pinyin, meanings, radical, stroke_count;
	`

	var updated entity.Ideogram

	err := r.db.QueryRow(
		ctx,
		query,
		ideog.Pinyin,
		ideog.Meanings,
		ideog.Radical,
		ideog.StrokeCount,
		character,
	).Scan(
		&updated.ID,
		&updated.Character,
		&updated.Pinyin,
		&updated.Meanings,
		&updated.Radical,
		&updated.StrokeCount,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, entity.ErrNotFound
		}

		return nil, fmt.Errorf(
			"erro ao atualizar ideograma no banco: %w",
			err,
		)
	}

	return &updated, nil
}
