package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	var data1 string
	var data2 string

	go func() {
		time.Sleep(10 * time.Second)
		channel1 <- "Hello"
	}()

	go func() {
		time.Sleep(5 * time.Second)
		channel2 <- "World!"
	}()

	//* select

	for len(data1) == 0 || len(data2) == 0 {
		select {
		case data1 = <-channel1:
			fmt.Println("Received data from channel1", data1)
		case data2 = <-channel2:
			fmt.Println("Received data from channel2", data2)
		default:
			fmt.Println("No data came")
		}
		time.Sleep(1 * time.Second)
	}

	// ! bu yapıda data 2 ye 5 saniye sonra veri
	// ! gelmesine rağmen data 1 blockladığı için
	// ! veriyi ekranda data 1 ile beraber yazar
	// ! bu gibi durumlarda "select" yapısı kullanılır.
	// data1 = <-channel1
	// data2 = <-channel2

	// fmt.Println("Data1 came : ", data1)
	// fmt.Println("Data2 came : ", data2)

}
