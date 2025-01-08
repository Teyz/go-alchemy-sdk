package alchemy

import (
	"fmt"
)

type AlchemyService struct {
	AlchemyUrl string
	Store      Alchemy
}

func NewAlchemyClient(apiKey string) (*AlchemyService, error) {
	return &AlchemyService{
		AlchemyUrl: fmt.Sprintf("https://solana-mainnet.g.alchemy.com/v2/%s", apiKey),
	}, nil
}
