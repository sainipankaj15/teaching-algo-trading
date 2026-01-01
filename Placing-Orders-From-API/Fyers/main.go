package main

import (
	"fmt"

	fyers "github.com/sainipankaj15/All-In-One-Broker/Fyers"
)

func main() {

	resp, err := fyers.PlaceMarketOrder("NSE:MOSMALL250-EQ", 1, fyers.TransactionSide.SELL, fyers.ProductType.INTRADAY, "XP03754")
	if err != nil {
		fmt.Println("Error placing order:", err)
		return
	}
	fmt.Println("Order placed successfully. Order ID:", resp)
}
