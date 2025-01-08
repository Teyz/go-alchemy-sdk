package solana

import rpc "github.com/teyz/go-alchemy-sdk/rpc"

type Solana interface {
	GetTransaction(signature string, commitment rpc.CommitmentType, maxSupportedTransactionVersion *uint64) (*rpc.GetTransactionResponse, error)
}
