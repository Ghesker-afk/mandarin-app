package entity

import "context"

type IdeogramRepository interface {
	FindByCharacter(ctx context.Context, char string) (*Ideogram, error)
	SearchByMeaning(ctx context.Context, meaning string) ([]Ideogram, error)
	Create(ctx context.Context, ideogram *Ideogram) (*Ideogram, error)
}
