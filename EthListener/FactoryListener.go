package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	//connect to local node
	cl, err := ethclient.Dial("http://127.0.0.1:7545")

	ctx := context.Background()

	addr := common.HexToAddress("0x60c3A0e473Af1fD6Ab33817d0562e7a67617F5fE")

	if err != nil {
		panic(err)
	}

	factory, err := NewWalletFactory(addr, cl)
	callOpts := &bind.CallOpts{Context: ctx, Pending: false}
	cost, err := factory.MintCost(callOpts)
	fmt.Println(cost)

	// Watch for a minted wallets events on Factory
	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}

	// Setup a channel for minted wallets
	walletsChannel := make(chan *WalletFactoryWalletMinted)

	var wg sync.WaitGroup

	// Start a goroutine which watches new mint events
	go func() {

		sub, err := factory.WatchWalletMinted(watchOpts, walletsChannel)
		if err != nil {
			panic(err)
		}
		fmt.Print("listening...")
		defer sub.Unsubscribe()

	}()

	/*

	   //process events in the channel

	*/

}
