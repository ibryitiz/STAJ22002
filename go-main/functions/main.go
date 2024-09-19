package main

import "fmt"

func main() {
	sum(10, 15)
	fmt.Println(sum2(10, 15))
	fmt.Println(sum3(10, 15))
	fmt.Println(sum4(10, 15, 34))
	fmt.Println(sum4(10, 15, 34, 45))
	fmt.Println(sum4(10, 15, 34, 45, 51))

}

// değer döndürmeyen
func sum(x int, y int) {
	fmt.Println(x + y)
}

//değer döndüren
func sum2(x int, y int) int {
	return x + y
}

// birden çok değer döndüren
func sum3(x int, y int) (int, int) {
	return x + y, x - y
}

// birden çok parametre alabilen

func sum4(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
