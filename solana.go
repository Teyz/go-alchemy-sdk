package alchemy

import (
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	rpc "github.com/teyz/go-alchemy-sdk/rpc"
)

func (a *Alchemy) GetAccountInfo(address string, encodingType rpc.EncodingType, commitment rpc.CommitmentType) (*rpc.GetAccountInfoResponse, error) {
	requestBody := &rpc.AlchemyRequest{
		JsonRPC: rpc.JsonRPCVersion2_0,
		Method:  rpc.MethodGetAccountInfo,
		Params: []interface{}{
			address,
			&rpc.GetAccountInfoParams{
				Commitment: commitment,
				Encoding:   encodingType,
			},
		},
		ID: 1,
	}

	response, err := a.makeCall(http.MethodPost, requestBody, &rpc.GetAccountInfoResponse{})
	if err != nil {
		return nil, err
	}

	parsedResponse, ok := response.(*rpc.GetAccountInfoResponse)
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	return parsedResponse, nil
}

func (a *Alchemy) GetBalance(address string) (*rpc.GetBalanceResponse, error) {
	requestBody := &rpc.AlchemyRequest{
		JsonRPC: rpc.JsonRPCVersion2_0,
		Method:  rpc.MethodGetBalance,
		Params: []interface{}{
			address,
		},
		ID: 1,
	}

	response, err := a.makeCall(http.MethodPost, requestBody, &rpc.GetBalanceResponse{})
	if err != nil {
		return nil, err
	}

	parsedResponse, ok := response.(*rpc.GetBalanceResponse)
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	return parsedResponse, nil
}

func (a *Alchemy) GetBlock(slotNumber uint64, commitment rpc.CommitmentType, maxSupportedTransactionVersion *uint64) (*rpc.GetBlockResponse, error) {
	requestBody := &rpc.AlchemyRequest{
		JsonRPC: rpc.JsonRPCVersion2_0,
		Method:  rpc.MethodGetBlock,
		Params: []interface{}{
			slotNumber,
			&rpc.GetBlockParams{
				Encoding:                       rpc.EncodingJson,
				TransactionDetails:             "full",
				Rewards:                        false,
				MaxSupportedTransactionVersion: maxSupportedTransactionVersion,
			},
		},
		ID: 1,
	}

	response, err := a.makeCall(http.MethodPost, requestBody, &rpc.GetBlockResponse{})
	if err != nil {
		fmt.Printf("error getting block: %v\n", err)
		return nil, err
	}

	parsedResponse, ok := response.(*rpc.GetBlockResponse)
	if !ok {
		fmt.Printf("invalid response type: %T", response)
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	spew.Dump(parsedResponse)

	return parsedResponse, nil
}

func (a *Alchemy) GetTransaction(signature string, commitment rpc.CommitmentType, maxSupportedTransactionVersion *uint64) (*rpc.GetTransactionResponse, error) {
	requestBody := &rpc.AlchemyRequest{
		JsonRPC: rpc.JsonRPCVersion2_0,
		Method:  rpc.MethodGetTransaction,
		Params: []interface{}{
			signature,
			&rpc.GetTransactionParams{
				Commitment:                     commitment,
				MaxSupportedTransactionVersion: maxSupportedTransactionVersion,
			},
		},
		ID: 1,
	}

	response, err := a.makeCall(http.MethodPost, requestBody, &rpc.GetTransactionResponse{})
	if err != nil {
		return nil, err
	}

	parsedResponse, ok := response.(*rpc.GetTransactionResponse)
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	spew.Dump(parsedResponse)

	return parsedResponse, nil
}