package main

import (
	"fmt"

	zerodha "github.com/sainipankaj15/All-In-One-Broker/Zerodha"
)

func main() {

	resp, err := zerodha.GetFunds("FC8173")
	if err != nil {
		fmt.Println("Error getting funds:", err)
		return
	}
	fmt.Printf("Funds: %+v\n", resp.Data.Equity.Net)
}
