package rpc

type GetBlockParams struct {
	Encoding                       EncodingType       `json:"encoding"`
	TransactionDetails             TransactionDetails `json:"transactionDetails"`
	Rewards                        bool               `json:"rewards"`
	Commitment                     CommitmentType     `json:"commitment"`
	MaxSupportedTransactionVersion *uint64            `json:"maxSupportedTransactionVersion"`
}

type GetBlockResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Result  struct {
		BlockHeight       uint64 `json:"blockHeight"`
		BlockTime         uint64 `json:"blockTime"`
		BlockHash         string `json:"blockhash"`
		ParentSlot        uint64 `json:"parentSlot"`
		PreviousBlockhash string `json:"previousBlockhash"`
		Transactions      []struct {
			Meta        Meta        `json:"meta"`
			Transaction Transaction `json:"transaction"`
		} `json:"transactions"`
	} `json:"result"`
	Error *AlchemyError `json:"error,omitempty"`
	ID    uint64        `json:"id"`
}
