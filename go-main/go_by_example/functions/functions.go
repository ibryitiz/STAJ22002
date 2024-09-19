package main

import "fmt"

func plus(a int, b int) int {
	toplam := a + b
	return toplam
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	result := plus(3, 5)
	fmt.Println("Result = ", result)

	plus := plusPlus(15, 25, 5)
	fmt.Println(plus)
}
