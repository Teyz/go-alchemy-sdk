package rpc

import (
	"encoding/json"

	"github.com/mr-tron/base58"
)

type Details struct {
	BlockTime   uint64       `json:"blockTime"`
	Meta        *Meta        `json:"meta"`
	Slot        uint64       `json:"slot"`
	Transaction *Transaction `json:"transaction"`
	Version     interface{}  `json:"version"`
}

type Transaction struct {
	Message struct {
		AccountKeys []string `json:"accountKeys"`
		Header      struct {
			NumReadonlySignedAccounts   uint64 `json:"numReadonlySignedAccounts"`
			NumReadonlyUnsignedAccounts uint64 `json:"numReadonlyUnsignedAccounts"`
			NumRequiredSignatures       uint64 `json:"numRequiredSignatures"`
		} `json:"header"`
		Instructions []struct {
			Accounts       []uint64 `json:"accounts"`
			Data           string   `json:"data"`
			ProgramIdIndex uint64   `json:"programIdIndex"`
			StackHeight    uint64   `json:"stackHeight"`
		} `json:"instructions"`
		RecentBlockhash string `json:"recentBlockhash"`
	} `json:"message"`
	Signatures []string `json:"signatures"`
}

type Meta struct {
	ComputeUnitsConsumed uint64               `json:"computeUnitsConsumed"`
	Err                  *interface{}         `json:"err,omitempty"`
	Fee                  uint64               `json:"fee"`
	InnerInstructions    []*InnerInstructions `json:"innerInstructions"`
	LoadedAddresses      LoadedAddresses      `json:"loadedAddresses"`
	LogMessages          []string             `json:"logMessages"`
	PostBalances         []uint64             `json:"postBalances"`
	PostTokenBalances    *[]interface{}       `json:"postTokenBalances"`
	PreBalances          []uint64             `json:"preBalances"`
	PreTokenBalances     *[]interface{}       `json:"preTokenBalances"`
	Rewards              *[]Rewards           `json:"rewards"`
	//DEPRECATED IN SOLANA
	Status interface{} `json:"status"`
}

type LoadedAddresses struct {
	Readonly []string `json:"readonly"`
	Writable []string `json:"writable"`
}

type InnerInstructions struct {
	Index        int `json:"index"`
	Instructions []struct {
		Accounts       []uint64 `json:"accounts"`
		Data           Base58   `json:"data"`
		ProgramIdIndex uint64   `json:"programIdIndex"`
		StackHeight    uint64   `json:"stackHeight"`
	}
}

type Rewards struct {
	Pubkey      string `json:"pubkey"`
	Lamports    uint64 `json:"lamports"`
	PostBalance uint64 `json:"postBalance"`
	RewardType  string `json:"rewardType"`
	Commission  uint64 `json:"commission"`
}

type Base58 []byte

func (t Base58) MarshalJSON() ([]byte, error) {
	return json.Marshal(base58.Encode(t))
}

func (t *Base58) UnmarshalJSON(data []byte) (err error) {
	var s string
	err = json.Unmarshal(data, &s)
	if err != nil {
		return
	}

	if s == "" {
		*t = []byte{}
		return nil
	}

	*t, err = base58.Decode(s)
	return
}

func (t Base58) String() string {
	return base58.Encode(t)
}
