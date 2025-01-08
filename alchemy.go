package alchemy

import (
	"context"
	"fmt"

	"github.com/teyz/go-alchemy-sdk/solana"
)

type AlchemyService struct {
	AlchemyUrl string
	Solana     solana.Solana
}

func NewAlchemyClient(ctx context.Context, apiKey string) (*AlchemyService, error) {
	alchemyUrl := fmt.Sprintf("https://solana-mainnet.g.alchemy.com/v2/%s", apiKey)

	solanaClient := solana.NewClient(ctx, alchemyUrl)
	return &AlchemyService{
		AlchemyUrl: alchemyUrl,
		Solana:     solanaClient,
	}, nil
}
