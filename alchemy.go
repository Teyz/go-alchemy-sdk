package alchemy

import (
	"fmt"
)

type Alchemy struct {
	AlchemyUrl string
}

func NewAlchemyClient(apiKey string) (*Alchemy, error) {
	return &Alchemy{
		AlchemyUrl: fmt.Sprintf("https://solana-mainnet.g.alchemy.com/v2/%s", apiKey),
	}, nil
}
