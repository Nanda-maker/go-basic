package main

import (
	"fmt"
	"go/cryptomasters/api"
	"sync"
)

// inside main function if there are multiple goroutines, we need to use waitgroup to wait for all goroutines to finish before exiting main function
//otherwise main function may exit before goroutines complete their execution since main function runs in its own goroutine and other goroutines run concurrently.

func main() {
	currencies := []string{"BTC", "BCH", "ETH"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		// func is lambda function or anonymous function
		// passing currency as parameter to avoid closure issue
		// each goroutine gets its own copy of currency variable
		// closure is when a function "closes over" variables from its surrounding scope
		// without passing currency as parameter, all goroutines may reference the same variable which can lead to unexpected results
		// by passing currency as parameter, each goroutine gets its own copy of the variable
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
