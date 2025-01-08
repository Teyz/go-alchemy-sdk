package main

import (
	"fmt"

	"github.com/teyz/go-alchemy-sdk"
	"github.com/teyz/go-alchemy-sdk/rpc"
)

func main() {
	alchemy, err := alchemy.NewAlchemyClient("K2UYNjx2z64VOHjdtDH7CITeVPiETvUP")
	if err != nil {
		fmt.Println("Error creating Alchemy client:", err)
		return
	}

	test := uint64(0)
	alchemy.Solana.GetTransaction("28XtqRGx5HNYe7qaZkLESD6MeGdHw2PApzBWs3dFF3bJtJ9rQfTLoMBRnpyU77cNt2tTD96bDtoLKc9FAhru4vyB", rpc.CommitmentFinalized, &test)
	//alchemy.GetAccountInfo("5eG7mnKDyi7zpb5wacK7K7GRy8pbgFKCBoL54WAe6kVj", rpc.EncodingBase58, rpc.CommitmentFinalized)
	//alchemy.GetBlock(307494832, rpc.CommitmentFinalized, &test)
	//alchemy.RequestAirdrop("5eG7mnKDyi7zpb5wacK7K7GRy8pbgFKCBoL54WAe6kVj", 1000000000, rpc.CommitmentFinalized)
}
