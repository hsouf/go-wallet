package listeners

import (
	"context"
	"fmt"
	"main/erc20"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	log "github.com/sirupsen/logrus"
	gomail "gopkg.in/mail.v2"
)

func ERC20Listener(wallets []common.Address, erc20Token string, endpoint string) {

	cl, err := ethclient.Dial(endpoint)

	ctx := context.Background()

	addr := common.HexToAddress(erc20Token)

	if err != nil {
		panic(err)
	}

	factory, err := erc20.NewErc20(addr, cl)

	watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
	backoffMax := 100 * time.Millisecond //generous backoff ^^

	// Setup a channel for deposited ether
	depositsChannel := make(chan *erc20.Erc20Transfer)

	//temporary array
	var from []common.Address
	from = make([]common.Address, 0)

	sub, err := factory.WatchTransfer(watchOpts, depositsChannel, from, wallets)
	if err != nil {
		panic(err)

	}
	fmt.Println("connection established, listenning for ether deposits on", erc20Token)

	for {
		select {

		case <-sub.Err():

			sub = event.Resubscribe(backoffMax, func(ctx context.Context) (event.Subscription, error) {
				subscription, err := factory.WatchTransfer(watchOpts, depositsChannel, from, wallets)

				if subscription != nil {
					log.WithFields(log.Fields{
						"socket":    endpoint,
						"timestamp": time.Now(),
						"listener":  "ERC20",
					}).Warn("Socket reconnected!")

				} else {

					log.WithFields(log.Fields{
						"timestamp": time.Now(),
						"socket":    endpoint,
					}).Warn("Socket disconnected, trying to reconnect...")
					//log reconnect message each 100 ms
					time.Sleep(backoffMax)

				}

				return subscription, err
			})

		case event := <-depositsChannel:

			go processDeposit(event, erc20Token)
		}
	}

}

func processDeposit(ch *erc20.Erc20Transfer, erc20 string) {

	event := ch

	fmt.Println("deposited amount", event.Value)
	fmt.Println("depositor:", event.From)
	fmt.Println("receiver: ", event.To)

	//fetch the linked email to wallet from DB

	//send email to user (if email is available ^^)

	m := gomail.NewMessage()
	m.SetHeader("From", "from@gmail.com")
	m.SetHeader("To", "to@example.com")
	m.SetHeader("Subject", "Gomail test subject")
	m.SetBody("text/plain", "This is Gomail test body")
	d := gomail.NewDialer("smtp.gmail.com", 587, "from@gmail.com", "<email_password>")

	if err := d.DialAndSend(m); err != nil {
		log.WithFields(log.Fields{
			"timestamp": time.Now(),
			"service":   "gmail",
			"listener":  "ERC20",
			"from":      "",
			"to":        "",
			"value":     "",
			"token":     erc20,
		}).Warn("Failed to send email", err)
	}

}
