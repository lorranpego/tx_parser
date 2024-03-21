package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"tx_parser/pkg/parser"
	_ "tx_parser/pkg/storage/memory" // mem storage
)

const (
	URL         = "https://cloudflare-eth.com"
	USDTADDRESS = "0xdac17f958d2ee523a2206206994597c13d831ec7"
)

func main() {
	p := parser.New(URL)
	in := bufio.NewScanner(os.Stdin)

	printOptions()
	for in.Scan() {
		switch in.Text() {
		case "exit":
			return
		case "block":
			block := p.GetCurrentBlock()
			log.Println(fmt.Sprintf("Last processed block: %d", block))
			printOptions()
		case "tx":
			log.Println(fmt.Sprintf("Enter address: Example %s", USDTADDRESS))
			if in.Scan() {
				address := in.Text()
				if len(address) <= 0 {
					log.Println("No address informed. Using standard USDT address.")
					address = USDTADDRESS
				}
				log.Println(fmt.Sprintf("Transactions for address: %s", address))

				transactions, err := p.GetTransactions(address)
				if err != nil {
					log.Println(fmt.Sprintf("Error when getting transactions for adddress: %s. Err: %v ", address, err))
				}
				if len(transactions) == 0 {
					log.Println("No address informed.")
				} else {
					marshal, _ := json.Marshal(transactions)
					log.Println(fmt.Sprintf("%s", string(marshal)))
				}
			}
			printOptions()
		case "sub":
			log.Println("Enter address: ")
			if in.Scan() {
				address := in.Text()
				if len(address) <= 0 {
					log.Println("No address informed. Using standard USDT address.")
					address = USDTADDRESS
				}
				if p.Subscribe(address) {
					log.Println(fmt.Sprintf("Address %s has been subscribed", address))
				} else {
					log.Println(fmt.Sprintf("Address %s has already been subscribed", address))
				}
			}

			printOptions()
		case "help":
			printOptions()
		default:
			log.Println("Unknown command. Type help for possible commands.")
		}
	}
}

func printOptions() {
	log.Println(`
	Commands:
		block: Returns last parsed block number
		sub: Subscribes to address in order to track it's transactions
		tx: Retreives list of address's transactions
		help: Prints menu for help
		exit: Exits application
	`)
}
