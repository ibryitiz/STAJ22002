package main

import "fmt"

func main() {

	//! foreach olmadan
	// var numbers = []int{1, 2, 3, 4, 5}

	// for index := 0; index < len(numbers); index++ {
	// 	fmt.Println(numbers[index])
	// }

	//? foreach yapısı
	var ages = []int{15, 25, 21, 28, 38}

	for index, value := range ages {
		fmt.Println(index, value)
	}

	//? index kullanmamak için
	for _, value := range ages {
		fmt.Println(value)
	}

	var language = "golang"

	for _, character := range language {
		fmt.Printf("Character %c\n", character)
	}

}
