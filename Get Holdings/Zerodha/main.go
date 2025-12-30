package main

import (
	"fmt"

	zerodha "github.com/sainipankaj15/All-In-One-Broker/Zerodha"
)

func main() {

	resp, err := zerodha.GetHoldings("FC8173")
	if err != nil {
		fmt.Println("Error getting holdings:", err)
		return
	}
	fmt.Printf("Holdings: %+v\n", resp)
}
