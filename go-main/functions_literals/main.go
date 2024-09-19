package main

import (
	"fmt"
	"reflect"
)

func main() {

	func() {
		fmt.Println("f1") // fonksiyon içinde fonksiyon
	}()

	func(x int, y int) {
		fmt.Println(x + y)
	}(5, 8)

	add := func(x int, y int) int {
		return x + y
	}

	fmt.Println(reflect.TypeOf(add))
	fmt.Println(add(8, 5))

	multiple := func(x int, y int) int {
		return x * y
	}

	v, y := calculator(8, 5, add, multiple)
	fmt.Println(v, y)

}

func calculator(x int, y int, f1 func(int, int) int, f2 func(int, int) int) (int, int) {
	return f1(x, y), f2(x, y)
}
