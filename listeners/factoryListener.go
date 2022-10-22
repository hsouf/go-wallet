package listeners

import (
	"context"
	"fmt"

	"main/walletFactory"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func FactoryListener(factoryAddress string) {

	//connect to local node
	cl, err := ethclient.Dial("ws://127.0.0.1:9999")

	ctx := context.Background()

	addr := common.HexToAddress(factoryAddress)

	if err != nil {
		panic(err)
	}

	factory, err := walletFactory.NewWalletFactory(addr, cl)
	callOpts := &bind.CallOpts{Context: ctx, Pending: false}
	cost, err := factory.MintCost(callOpts)
	fmt.Println("Factory: mint cost", cost, "wei")

	// Watch for a minted wallets events on Factory
	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}

	// Setup a channel for minted wallets
	walletsChannel := make(chan *walletFactory.WalletFactoryWalletMinted)

	//var wg sync.WaitGroup
	var socketConnected = false
	for {

		//connect to socket
		if socketConnected == false {
			sub, err := factory.WatchWalletMinted(watchOpts, walletsChannel)
			if err != nil {
				panic(err)

			}
			fmt.Println("connection established, listenning for mint events on", factoryAddress)
			defer func() {
				sub.Unsubscribe()
				socketConnected = false

			}()

		}
		socketConnected = true

		event := <-walletsChannel
		go addWalletToDb(event)

	}

}

func addWalletToDb(ch *walletFactory.WalletFactoryWalletMinted) {

	event := ch

	fmt.Println("received admin:", event.Admin)

	fmt.Println("received wallet address:", event.Wallet)

	/*

	   handle admin + wallet address

	*/
	fmt.Print("Finished adding wallet to db \n")

}
