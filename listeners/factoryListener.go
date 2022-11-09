package listeners

import (
	"context"
	"fmt"

	"main/walletFactory"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	log "github.com/sirupsen/logrus"
)

func FactoryListener(factoryAddress string, endpoint string) { //connect to local node

	cl, err := ethclient.Dial(endpoint)

	ctx := context.Background()

	addr := common.HexToAddress(factoryAddress)

	if err != nil {
		panic(err)
	}

	////////////////////////////////////////////////////////////////////////////////////////
	factory, err := walletFactory.NewWalletFactory(addr, cl)
	callOpts := &bind.CallOpts{Context: ctx, Pending: false}
	cost, err := factory.MintCost(callOpts)
	fmt.Println("Factory: mint cost", cost, "wei")

	// Watch for a minted wallets events on Factory
	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}

	fmt.Println("connecting to socket.....")
	backoffMax := 100 * time.Millisecond //generous backoff ^^
	walletsChannel := make(chan *walletFactory.WalletFactoryWalletMinted)

	sub, err := factory.WatchWalletMinted(watchOpts, walletsChannel)

	if err != nil {
		//first connection needs to succeed
		panic(err)
	}

	log.WithFields(log.Fields{
		"factory": factoryAddress,
	}).Info("connection established, listenning for mint events on")

	defer func() {
		sub.Unsubscribe()

	}()

	for {
		select {

		case <-sub.Err():

			sub = event.Resubscribe(backoffMax, func(ctx context.Context) (event.Subscription, error) {
				subscription, err := factory.WatchWalletMinted(watchOpts, walletsChannel)

				if subscription != nil {
					log.WithFields(log.Fields{
						"socket":    endpoint,
						"timestamp": time.Now(),
						"listener":  "Factory",
					}).Warn("Socket reconnected!")

				} else {

					log.WithFields(log.Fields{
						"timestam   p": time.Now(),
						"socket": endpoint,
					}).Warn("Socket disconnected, trying to reconnect...")
					//log reconnect message each 100 ms
					time.Sleep(backoffMax)

				}

				return subscription, err
			})

		case event := <-walletsChannel:

			go addWalletToDb(event)
		}
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
