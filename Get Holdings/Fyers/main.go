package main

import (
	"fmt"

	fyers "github.com/sainipankaj15/All-In-One-Broker/Fyers"
)

func main() {

	resp, err := fyers.GetHoldings("XP03754")
	if err != nil {
		fmt.Println("Error getting Holdings:", err)
		return
	}
	fmt.Printf("Holdings: %+v\n", resp)
}
