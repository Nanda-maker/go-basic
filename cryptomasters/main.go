package main

import (
	"fmt"
	"go/cryptomasters/api"
	"sync"
)

func main() {
	currencies := []string{"BTC", "BCH", "ETH"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getcurrencyData(currencyCode)
			wg.Done()
		}(currency)

	}
	wg.Wait()
}

func getcurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
