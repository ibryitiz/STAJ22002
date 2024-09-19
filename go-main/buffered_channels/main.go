package main

import (
	"fmt"
	"time"
)

func main() {
	//myChannel := make(chan int, 2)

	// myChannel <- 1
	// myChannel <- 2

	// fmt.Println(<-myChannel) // ilk gönderilen ilk okunur = 1

	// myChannel <- 3
	// fmt.Println(<-myChannel) // Sıra 2 de çıktı : 2

	messages := make(chan string, 2)

	go func() {
		data1 := <-messages
		fmt.Println("First : ", data1)
		data2 := <-messages
		fmt.Println("Second : ", data2)
	}()

	messages <- "Hello"
	messages <- "World"
	messages <- "World2"

	time.Sleep(1 * time.Second)

	fmt.Println("End of the main")

}
