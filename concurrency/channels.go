package concurrency

import (
	"fmt"
)

func sendValue(c chan string) {
	c <- "Hello World!"
}

func ChannelsMain() {

	fmt.Println("Concurrency with Channels")

	values := make(chan string, 2) // Initialized buffered channel with second parameter that allows send command to be non-blocking until the channel buffer is full.
	defer close(values)


	//go sendValue(values)
	//go sendValue(values) 
	// Buffer channel allowed above commands to run in main routine. Otherwise it would have blocked the main routine to deadlock.

	sendValue(values)
	sendValue(values) 

	value := <-values
	value2 := <-values
	fmt.Println(value)
	fmt.Println(value2)

}