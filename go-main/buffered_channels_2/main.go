package main

import (
	"fmt"
	"time"
)

func main() {
	bufferedChannels := make(chan int, 4)

	go func() {
		for i := 1; i <= 10; i++ {
			bufferedChannels <- i
			fmt.Println("Send data : ", i)
		}
		close(bufferedChannels)
	}()

	time.Sleep(3 * time.Second)

	for data := range bufferedChannels {
		fmt.Println("Received data : ", data)
		time.Sleep(1 * time.Second)
	}
}
