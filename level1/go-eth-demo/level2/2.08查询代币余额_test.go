package main

import (
	// 	"fmt"
	// 	"log"
	// 	"math"
	// 	"math/big"
	"testing"
	// token "github.com/local/go-eth-demo/erc20" // for demo
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/ethclient"
)

func TestQueryTokenBalance(t *testing.T) {

	// 	// client, err := ethclient.Dial("https://sepolia.infura.io/v3/49bd5a6c7f1a42b8bae5c825f15578a1")
	// 	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	//	 // Golem (GNT) Address
	//	 tokenAddress := common.HexToAddress("0xfadea654ea83c00e5003d2ea15c59830b65471c0")
	//	 instance, err := token.NewToken(tokenAddress, client)
	//	 if err != nil {
	//			 log.Fatal(err)
	//	 }
	//	 address := common.HexToAddress("0x25836239F7b632635F815689389C537133248edb")
	//	 bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	//	 if err != nil {
	//			 log.Fatal(err)
	//	 }
	//	 name, err := instance.Name(&bind.CallOpts{})
	//	 if err != nil {
	//			 log.Fatal(err)
	//	 }
	//	 symbol, err := instance.Symbol(&bind.CallOpts{})
	//	 if err != nil {
	//			 log.Fatal(err)
	//	 }
	//	 decimals, err := instance.Decimals(&bind.CallOpts{})
	//	 if err != nil {
	//			 log.Fatal(err)
	//	 }
	//	 fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	//	 fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	//	 fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	//	 fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"
	//	 fbal := new(big.Float)
	//	 fbal.SetString(bal.String())
	//	 value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	//	 fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
