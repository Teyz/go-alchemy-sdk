package alchemy

import (
	"fmt"

	"github.com/teyz/go-alchemy-sdk/solana"
)

type AlchemyService struct {
	AlchemyUrl string
	Solana     solana.Solana
}

func NewAlchemyClient(apiKey string) (*AlchemyService, error) {
	return &AlchemyService{
		AlchemyUrl: fmt.Sprintf("https://solana-mainnet.g.alchemy.com/v2/%s", apiKey),
	}, nil
}
