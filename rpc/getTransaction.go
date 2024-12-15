package rpc

type GetTransactionParams struct {
	Commitment                     CommitmentType `json:"commitment"`
	MaxSupportedTransactionVersion *uint64        `json:"maxSupportedTransactionVersion"`
}

type GetTransactionResponse struct {
	JsonRPC string        `json:"jsonrpc"`
	Result  *Details      `json:"result"`
	Error   *AlchemyError `json:"error,omitempty"`
	ID      uint64        `json:"id"`
}
