package rpc

type AlchemyRequest struct {
	JsonRPC JsonRPCVersion `json:"jsonrpc"`
	Method  Method         `json:"method"`
	Params  []interface{}  `json:"params"`
	ID      int            `json:"id"`
}
