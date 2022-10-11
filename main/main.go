package main

import (
	"fmt"
	"main/listeners"
)

func main() {

	var wallets map[string]int
	wallets = make(map[string]int)
	wallets["soufiane"] = 900

	fmt.Println(wallets["soufiane"])

	listeners.FactoryListener("0xBb7Db47A8bE34246B6F29078C99523fd910533EB")

}
