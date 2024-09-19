package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//? Normalde bu yönetimi goRoutine leri Channels'lar yapar
//? ama onlar yokken bu şekilde yapılır

func main() {

	wg.Add(1)

	go printX()
	wg.Wait()
	printY()

	//time.Sleep(time.Second) //! kötü yöntem

}

func printX() {
	for i := 0; i < 1000; i++ {
		fmt.Print("X")
	}
	wg.Done()
}

func printY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
}
