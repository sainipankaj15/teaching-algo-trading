package main

import (
	"fmt"

	zerodha "github.com/sainipankaj15/All-In-One-Broker/Zerodha"
)

func main() {

	resp, err := zerodha.PlaceMarketOrder(zerodha.Exchange.NSE, "SBIN", "1", zerodha.OrderType.MARKET, zerodha.TransactionSide.SELL, zerodha.ProductType.INTRADAY, "FC8173")
	if err != nil {
		fmt.Println("Error placing order:", err)
		return
	}
	fmt.Println("Order placed successfully. Order ID:", resp)
}
