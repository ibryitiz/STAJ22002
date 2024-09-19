package main

import "fmt"

func main() {

	// var age int
	// age = 25

	// if age < 18 {
	// 	fmt.Println("ehliyet alamazsınz")
	// } else {
	// 	fmt.Println("ehliyet alabilirsiniz")

	// }

	a := 40
	b := 20
	c := 30

	if a > b && a > c {
		fmt.Printf("a = %d en büyük sayıdır.", a)
	} else if b > a && b > c {
		fmt.Printf("b = %d en büyük sayıdır.", b)
	} else {
		fmt.Printf("c = %d en büyük sayıdır.", c)
	}

}
