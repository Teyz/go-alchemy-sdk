package solana

import (
	"context"
)

type solanaClient struct {
	AlchemyUrl string
}

func NewClient(ctx context.Context, alchemyUrl string) Solana {
	return &solanaClient{
		AlchemyUrl: alchemyUrl,
	}
}
