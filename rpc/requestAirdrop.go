package rpc

type RequestAirdropRequest struct {
	Lamports   uint64         `json:"lamports"`
	Commitment CommitmentType `json:"commitment"`
}

type RequestAirdropResponse struct {
	JsonRPC string        `json:"jsonrpc"`
	Result  string        `json:"result"`
	Error   *AlchemyError `json:"error,omitempty"`
	ID      uint64        `json:"id"`
}
