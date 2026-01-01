package main

import (
	"fmt"

	fyers "github.com/sainipankaj15/All-In-One-Broker/Fyers"
)

func main() {

	resp, err := fyers.GetPositions("XP03754")
	if err != nil {
		fmt.Println("Error getting positions:", err)
		return
	}
	fmt.Printf("Positions: %+v\n", resp.Overall.PLRealized)
}
