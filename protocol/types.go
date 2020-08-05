package protocol

// Structure of the Bitcoin Block Header
type BlockHeader struct {
	version           int32
	previousBlockHash [32]byte
	merkleRoot        [32]byte
	timestamp         uint32
	bits              uint32
	nonce             uint32
}
