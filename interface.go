package alchemy

import rpc "github.com/teyz/go-alchemy-sdk/rpc"

type Alchemy interface {
	GetTransaction(signature string, commitment rpc.CommitmentType, maxSupportedTransactionVersion *uint64) (*rpc.GetTransactionResponse, error)
}
