package rpc

type AlchemyError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
