package listeners

import (
	"context"
	"fmt"
	"main/erc20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ERC20Listener(walletAddress string) {

	//connect to local node
	cl, err := ethclient.Dial("ws://127.0.0.1:9999")

	ctx := context.Background()

	addr := common.HexToAddress(walletAddress)

	if err != nil {
		panic(err)
	}

	factory, err := erc20.NewErc20(addr, cl)

	// Watch for a minted wallets events on Factory
	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}

	// Setup a channel for deposited ether
	depositsChannel := make(chan *erc20.Erc20Transfer)

	//temporary array
	var from []common.Address
	from = make([]common.Address, 1)

	var to []common.Address
	to = make([]common.Address, 1)

	//var wg sync.WaitGroup
	var socketConnected = false
	for {

		//connect to socket
		if socketConnected == false {
			sub, err := factory.WatchTransfer(watchOpts, depositsChannel, from, to)
			if err != nil {
				panic(err)

			}
			fmt.Println("connection established, listenning for ether deposits on", walletAddress)

			defer func() {
				sub.Unsubscribe()
				socketConnected = false

			}()

		}
		socketConnected = true

		event := <-depositsChannel
		go addErc20TransferToDb(event)

	}

}

func addErc20TransferToDb(ch *erc20.Erc20Transfer) {

	event := ch

	fmt.Println("deposited amount", event.Value)
	fmt.Println("depositor:", event.From)
	fmt.Println("receiver: ", event.To)
	/*
	   @TODO: add deposit data to db
	*/

}
