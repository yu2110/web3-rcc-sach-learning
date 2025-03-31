package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var StoreABI = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

func TestContractEvent(t *testing.T) {

	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("0x2958d15bc5b64b11Ec65e623Ac50C198519f8742")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(StoreABI)))
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex())
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}
			err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(common.Bytes2Hex(event.Key[:]))
			fmt.Println(common.Bytes2Hex(event.Value[:]))
			var topics []string
			for i := range vLog.Topics {
				topics = append(topics, vLog.Topics[i].Hex())
			}
			fmt.Println("topics[0]=", topics[0])
			if len(topics) > 1 {
				fmt.Println("index topic:", topics[1:])
			}
		}
	}
}
