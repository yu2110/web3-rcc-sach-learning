package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"testing"

	// "github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestQueryTransactions(t *testing.T) {

	// client, err := ethclient.Dial("https://sepolia.infura.io/v3/49bd5a6c7f1a42b8bae5c825f15578a1")
	client, err := ethclient.Dial("https://ethereum.publicnode.com/?sepolia")

	if err != nil {
		log.Fatal(err)
	}

	// chainID, err := client.ChainID(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	blockNumber := big.NewInt(5171744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(tx.Value().String())    // 100000000000000000
		fmt.Println(tx.Gas())               // 21000
		fmt.Println(tx.GasPrice().Uint64()) // 100000000000
		fmt.Println(tx.Nonce())             // 245132
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587

		// if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
		// 	fmt.Println("sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
		// } else {
		// 	log.Fatal(err)
		// }

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
		fmt.Println(receipt.Logs)   // []
		break
	}

	blockHash := common.HexToHash("0x3aa9448d513c60fe63b04c4cbc8641d5dbb4fa943ab9efc2c4294cde3a305d25")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		break
	}

	// txHash := common.HexToHash("0xe8d7aa873a706c43413c2b0151e22ae1f717fb4366c4b3a4796c6619d53b4c7b")
	// tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(isPending)
	// fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5.Println(isPending)       // false
}
