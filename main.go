package main

import (
	"main/listeners"

	"github.com/ethereum/go-ethereum/common"
)

func main() {

	ethEndpoint := "ws://127.0.0.1:9999"
	erc20Token := ""

	//temp array, all wallets to listen to should be fetched every 10minutes from DB
	var wallets []common.Address
	wallets = make([]common.Address, 1)

	wallets[0] = common.HexToAddress("0xBb7Db47A8bE34246B6F29078C99523fd910533EB")

	//listeners.FactoryListener("0xBb7Db47A8bE34246B6F29078C99523fd910533EB", ethEndpoint)
	listeners.ERC20Listener(wallets, erc20Token, ethEndpoint)

}
