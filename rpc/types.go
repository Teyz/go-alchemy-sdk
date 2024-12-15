package rpc

type JsonRPCVersion string

const (
	JsonRPCVersion2_0 JsonRPCVersion = "2.0"
	JsonRPCVersion1_0 JsonRPCVersion = "1.0"
)

type Method string

const (
	MethodGetAccountInfo Method = "getAccountInfo"
	MethodGetTransaction Method = "getTransaction"
	MethodGetBalance     Method = "getBalance"
	MethodGetBlock       Method = "getBlock"
)

type CommitmentType string

const (
	CommitmentFinalized CommitmentType = "finalized"
	CommitmentConfirmed CommitmentType = "confirmed"
	CommitmentProcessed CommitmentType = "processed"
)

type EncodingType string

const (
	EncodingBase64     EncodingType = "base64"
	EncodingBase58     EncodingType = "base58"
	EncodingBase64Zstd EncodingType = "base64+zstd"
	EncodingJsonParsed EncodingType = "jsonParsed"
	EncodingJson       EncodingType = "json"
)
