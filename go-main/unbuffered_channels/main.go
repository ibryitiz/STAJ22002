/*package main

import (
	"fmt"
)

func main() {
	myChannel := make(chan string)

	// go func() {
	// 	myChannels <- "Hello World!"
	// }()

	// // Blocking
	// message, isOpen := <-myChannels

	// fmt.Println(message, isOpen)

	//? -------------------

	done := make(chan bool)

	go func() {
		message := <-myChannel
		fmt.Println(message)
		done <- true
	}()

	go func() {
		myChannel <- "Hello World!"
	}()

	<-done

	fmt.Println("End of the main")

}*/

package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			myChannel <- i
			fmt.Println("Send data : ", i)
			time.Sleep(1 * time.Second)
		}
		close(myChannel) // channel kaptama
		//! eğer kapatmazsak deadlock! yeriz
	}()

	for {
		data, isOpen := <-myChannel

		if isOpen == false {
			break
		}

		fmt.Println("Received data : ", data)
	}
}
