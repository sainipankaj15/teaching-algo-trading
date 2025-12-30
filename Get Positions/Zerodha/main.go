package main

import (
	"fmt"

	zerodha "github.com/sainipankaj15/All-In-One-Broker/Zerodha"
)

func main() {

	resp, err := zerodha.GetPositions("FC8173")
	if err != nil {
		fmt.Println("Error getting positions:", err)
		return
	}
	fmt.Printf("Positions: %+v\n", resp)
}
