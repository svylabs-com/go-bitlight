package protocol

const magic = "0xD9B4BEF9"

type VarInt struct {
}

// Structure of the Bitcoin Block Header
type BlockHeader struct {
	version           int32
	previousBlockHash [32]byte
	merkleRoot        [32]byte
	timestamp         uint32
	bits              uint32
	nonce             uint32
	transactionCount  VarInt
}

type Headers struct {
	count   VarInt
	headers []BlockHeader
}

type Message struct {
	magic        uint32
	command      [12]byte
	paylodLength uint32
	checksum     uint32
	payload      []byte
}

type VersionMessage struct {
}

type GetHeadersMessage struct {
	version          uint32
	hashCount        VarInt
	blockLocatorHash [32]byte
	hashStop         [32]byte
}
