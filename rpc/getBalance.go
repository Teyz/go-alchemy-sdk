package rpc

type GetBalanceResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Result  struct {
		Context struct {
			ApiVersion string `json:"apiVersion"`
			Slot       uint64 `json:"slot"`
		} `json:"context"`
		Value uint64 `json:"value"`
	} `json:"result"`
	Error *AlchemyError `json:"error,omitempty"`
	ID    uint64        `json:"id"`
}
