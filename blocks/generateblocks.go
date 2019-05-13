package blockchainnetwork

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	Data      string
	Timestamp string
	PrevHash  string
	Hash      string
}

var GenesisBlock = Block{
	Index:     1,
	Data:      "genesisblock",
	Timestamp: time.Now().String(),
	PrevHash:  "",
	Hash:      "",
}

var Blockchain []Block

//Uses SHA256 algorithm to generate hash value of a Block.
func GenerateHash(b Block) (hashValue string) {
	val := string(b.Index) + b.Data + b.Timestamp + b.PrevHash
	h := sha256.New()
	h.Write([]byte(val))
	hashValue = hex.EncodeToString(h.Sum(nil))
	return hashValue
}

//Generate a new Block.
func GenerateBlock(oldBlock Block, newData string) Block {
	newBlock := Block{
		Index:     oldBlock.Index + 1,
		Data:      newData,
		Timestamp: time.Now().String(),
		PrevHash:  oldBlock.Hash,
	}

	newBlock.Hash = GenerateHash(newBlock)

	return newBlock
}

//Check if the new Block generated is valid.
func CheckBlock(oldBlock, newBlock Block) bool {
	if newBlock.PrevHash != oldBlock.Hash {
		return false
	}

	if newBlock.Index != oldBlock.Index+1 {
		return false
	}

	if GenerateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true

}

//Chaining the Blocks
func GenerateChain(newBlocks []Block) {

	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}