package listeners

import (
	"context"
	"fmt"
	"main/wallet"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func EthListener(walletAddress common.Address, endpoint string) {

	//connect to local node
	cl, err := ethclient.Dial(endpoint)

	ctx := context.Background()

	if err != nil {
		panic(err)
	}

	factory, err := wallet.NewWallet(walletAddress, cl)

	// Watch for a minted wallets events on Factory
	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}

	// Setup a channel for deposited ether
	depositsChannel := make(chan *wallet.WalletEtherDeposited)

	//temporary array
	var wallet []common.Address
	wallet = make([]common.Address, 1)
	wallet[0] = walletAddress

	//var wg sync.WaitGroup
	var socketConnected = false
	for {

		//connect to socket
		if socketConnected == false {
			sub, err := factory.WatchEtherDeposited(watchOpts, depositsChannel, wallet)
			if err != nil {
				panic(err)

			}

			defer func() {
				sub.Unsubscribe()
				socketConnected = false

			}()

		}
		socketConnected = true

		event := <-depositsChannel
		go addEthDepositToDb(event)

	}

}

func addEthDepositToDb(ch *wallet.WalletEtherDeposited) {

	event := ch

	fmt.Println("deposited amount", event.Amount)
	fmt.Println("depositor:", event.Depositor)
	fmt.Println("receiver: ", event.Wallet)
	/*
	   @TODO: add deposit data to db
	*/

}
