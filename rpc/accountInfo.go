package rpc

type GetAccountInfoParams struct {
	Commitment CommitmentType `json:"commitment"`
	Encoding   EncodingType   `json:"encoding"`
}

type GetAccountInfoResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Result  struct {
		Context struct {
			ApiVersion string `json:"apiVersion"`
			Slot       uint64 `json:"slot"`
		} `json:"context"`
		Value struct {
			Data       []string `json:"data"`
			Executable bool     `json:"executable"`
			Lamports   uint64   `json:"lamports"`
			Owner      string   `json:"owner"`
			RentEpoch  uint64   `json:"rentEpoch"`
			Space      uint64   `json:"space"`
		} `json:"value"`
	} `json:"result"`
	Error *AlchemyError `json:"error,omitempty"`
	ID    uint64        `json:"id"`
}
