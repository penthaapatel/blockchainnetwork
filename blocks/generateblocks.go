package blockchainnetwork

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

//Block A simplified version of a block
type Block struct {
	Index     int    `json:"index"`
	Data      string `json:"data"`
	Timestamp string `json:"timestamp"`
	PrevHash  string `json:"prevhash"`
	Hash      string `json:"hash"`
}

//GenesisBlock First block of the blockchain
var GenesisBlock = Block{
	Index:     0,
	Data:      "genesisblock",
	Timestamp: time.Now().String(),
	PrevHash:  "",
	Hash:      "",
}

//Blockchain An array of blocks
var Blockchain []Block

//GenerateHash Uses SHA256 algorithm to generate hash value of a Block.
func GenerateHash(b Block) (hashValue string) {
	val := string(b.Index) + b.Data + b.Timestamp + b.PrevHash
	h := sha256.New()
	h.Write([]byte(val))
	hashValue = hex.EncodeToString(h.Sum(nil))
	return hashValue
}

//GenerateBlock Generate a new Block.
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

//CheckBlock Check if the new Block generated is valid.
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

//GenerateChain Chaining the Blocks
func GenerateChain(newBlocks []Block) {

	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

//OutputJSON Pretty print the blocks
func OutputJSON(b Block) (string, error) {
	jsonOutput, err := json.MarshalIndent(b, "", "   ")
	if err != nil {
		return "", err
	}
	return string(jsonOutput), nil
}
