package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	//connect to local node
	cl, err := ethclient.Dial("ws://127.0.0.1:9999")

	ctx := context.Background()

	addr := common.HexToAddress("0x0E8340E63fF0BE0528eA1A70E3dedAfF356Fbe44")

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

	//var wg sync.WaitGroup
	var socketConnected = false
	for {

		//connect to socket
		if socketConnected == false {
			sub, err := factory.WatchWalletMinted(watchOpts, walletsChannel)
			if err != nil {
				panic(err)

			}

			defer sub.Unsubscribe()

		}
		socketConnected = true

		event := <-walletsChannel
		go addWalletToDb(event)

	}

}

func addWalletToDb(ch *WalletFactoryWalletMinted) {

	event := ch

	fmt.Println("received admin:", event.Admin)
	fmt.Println("received wallet address:", event.Wallet)

	/*

	   handle admin + wallet address

	*/
	fmt.Print("Finished adding wallet to db ")

}
