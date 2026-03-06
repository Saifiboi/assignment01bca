package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
)

type Block struct {
	transaction   string
	nonce        int
	prevHash     string

}

func NewBlock(transaction string, nonce int, prevHash string) *Block {
	return &Block{transaction, nonce, prevHash}
}

func ListBlocks(Blockchain []Block) {
	width := 82
	border := "+" + repeatChar("-", width) + "+"
	for i, block := range Blockchain {
		fmt.Println(border)
		fmt.Print("| BLOCK-", i,repeatChar(" ",width-8),"|\n")
		fmt.Println("|" + repeatChar("-", width) + "|")
		contentTx := "  Transaction : " + block.transaction
		fmt.Println("| " + contentTx + repeatChar(" ", width-len(contentTx)-2) + " |")
		contentNonce := "  Nonce       : " + fmt.Sprint(block.nonce)
		fmt.Println("| " + contentNonce + repeatChar(" ", width-len(contentNonce)-2) + " |")
		contentHash := "  Prev Hash   : " + block.prevHash
		fmt.Println("| " + contentHash + repeatChar(" ", width-len(contentHash)-2) + " |")	
		fmt.Println(border)
		if i < len(Blockchain)-1 {
			fmt.Print(repeatChar(" ", width/2), "|\n")
			fmt.Print(repeatChar(" ", width/2), "v\n")
		}
	}
}

func repeatChar(ch string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += ch
	}
	return result
}

func ChangeBlock(Blockchain []Block, index int, newTransaction string) {
	if index >= 0 && index < len(Blockchain) {
		Blockchain[index].transaction = newTransaction
	}
}

func CalculateHash(block Block) string {
	record := strconv.Itoa(block.nonce) + block.transaction + block.prevHash
	hash := sha256.New()
	hash.Write([]byte(record))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func VerifyChain(Blockchain []Block) bool {
	for i := 1; i < len(Blockchain); i++ {
		if Blockchain[i].prevHash != CalculateHash(Blockchain[i-1]) {
			return false
		}	
	}
	return true
}

func AddBlock(Blockchain []Block, transaction string) []Block {
	prevHash := CalculateHash(Blockchain[len(Blockchain)-1])
	nonce := rand.Intn(1000)
	block := NewBlock(transaction, nonce, prevHash)
	return append(Blockchain, *block)
}