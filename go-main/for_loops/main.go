package main

import "fmt"

func main() {
	// index := 1

	// for index <= 10 {
	// 	fmt.Println(index)
	// 	index = index + 1
	// }

	// toplam := 0

	// for i := 0; i <= 10; i++ {
	// 	toplam = toplam + i
	// }
	// fmt.Printf("1 den 10 a kadar olan sayıların toplamı = %d", toplam)

	// index := 0

	// var names = [3]string{"ibrahim", "ali", "murselat"}

	// for index < 3 {
	// 	fmt.Println(names[index])
	// 	fmt.Println("--------")
	// 	index++
	// }

	// index := 0

	// var names = [3]string{"ibrahim", "ali", "murselat"}

	// for index < len(names) {
	// 	fmt.Println(names[index])
	// 	fmt.Println("--------")
	// 	index++
	// }

	// index := 0

	// var names = [3]string{"ibrahim", "ali", "murselat"}

	// for index < len(names) {
	// 	var name string = names[index]
	// 	fmt.Println(name)
	// 	index++
	// }

	for index := 0; index <= 10; index++ {
		if index == 3 {
			break
		}
		fmt.Println(index)
	}

	fmt.Println("--------")

	for i := 0; i < 10; i++ {
		if i == 2 || i == 5 {
			continue
		}
		fmt.Println(i)
	}

}
