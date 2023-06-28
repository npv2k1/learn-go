// // connect go to ethereum

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	// "math/big"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func main() {
//     client, err := ethclient.Dial("HTTP://0.0.0.0:7545")
//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Println("we have a connection")
//     _ = client // we'll use this in the upcoming sections

// 		account := common.HexToAddress("0x5afC5c96620F1389afAD3722177799DaeE12a868")
// 		balance, err := client.BalanceAt(context.Background(), account, nil)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Println(balance)
// }