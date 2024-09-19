package main

import "fmt"

func fac(x float32) (y float32) {

	if x > 0 {
		y = x * fac(x-1)

	} else {
		y = 1
	}
	return
}

func main() {
	fact := fac(5)
	fmt.Println("5 sayısının faktoriyeli =", fact)
}
