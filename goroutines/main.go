package main

import (
	"fmt"
	"time"
)

// Threads - different lines of execution running at the same time.Goroutine will do that
func printMessage(text string, channel chan string){
	for i :=0; i<5;i++ {
		fmt.Println(text)
		time.Sleep( 800 * time.Millisecond)
	}
	channel <- "Done!"
	
}

func main() { // main goroutine
    
	var channel chan string
	go printMessage("Go is great",channel)
	response := <-channel
	fmt.Println(response)
//go printMessage("waiting") 
//printMessage("We miss classes")   
// for {}
}

// GoRoutine is simple way that go has to create threads and open threads
// Channel is the one way that we have to send and receive data between goroutines.