package main

import (
	"fmt"

	fyers "github.com/sainipankaj15/All-In-One-Broker/Fyers"
)

func main() {

	resp, err := fyers.GetFunds("XP03754")
	if err != nil {
		fmt.Println("Error getting Funds:", err)
		return
	}
	fmt.Printf("Funds: %+v\n", resp.FundLimit[9].EquityAmount)
}
