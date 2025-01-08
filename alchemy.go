package alchemy

import (
	"fmt"
)

type alchemy struct {
	AlchemyUrl string
}

func NewAlchemyClient(apiKey string) (*alchemy, error) {
	return &alchemy{
		AlchemyUrl: fmt.Sprintf("https://solana-mainnet.g.alchemy.com/v2/%s", apiKey),
	}, nil
}
