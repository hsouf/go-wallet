package listeners

import (
	"context"
	"fmt"
	"main/wallet"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func EthListener(walletAddress string) {

	//connect to local node
	cl, err := ethclient.Dial("ws://127.0.0.1:9999")

	ctx := context.Background()

	addr := common.HexToAddress(walletAddress)

	if err != nil {
		panic(err)
	}

	factory, err := wallet.NewWallet(addr, cl)

	// Watch for a minted wallets events on Factory
	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}

	// Setup a channel for deposited ether
	depositsChannel := make(chan *wallet.WalletEtherDeposited)

	//temporary array
	var wallets []common.Address
	wallets = make([]common.Address, 1)
	wallets[0] = addr

	//var wg sync.WaitGroup
	var socketConnected = false
	for {

		//connect to socket
		if socketConnected == false {
			sub, err := factory.WatchEtherDeposited(watchOpts, depositsChannel, wallets)
			if err != nil {
				panic(err)

			}
			fmt.Println("connection established, listenning for ether deposits on", walletAddress)
			defer sub.Unsubscribe()

		}
		socketConnected = true

		event := <-depositsChannel
		go addTransactionToDb(event)

	}

}

func addTransactionToDb(ch *wallet.WalletEtherDeposited) {

	event := ch

	fmt.Println("deposited amount", event.Amount)
	fmt.Println("depositor:", event.Depositor)
	fmt.Println("receiver: ", event.Wallet)
	/*
	   @TODO: add deposit data to db
	*/

}
