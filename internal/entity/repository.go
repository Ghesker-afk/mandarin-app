package entity

import "context"

type IdeogramRepository interface {
	FindAll(ctx context.Context) ([]Ideogram, error)
	FindByCharacter(ctx context.Context, char string) (*Ideogram, error)
	SearchByMeaning(ctx context.Context, meaning string) ([]Ideogram, error)
	Create(ctx context.Context, ideogram *Ideogram) (*Ideogram, error)
	Update(ctx context.Context, character string, ideogram *Ideogram) (*Ideogram, error)
}
