package models

type RPCBlock struct {
	BlockNumber int64
	// agnostic blob of data that is the block
	Payload []byte
}