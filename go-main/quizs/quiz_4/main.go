// 0-10 arasındaki sayıları çift ve tek diye yazdır

// package main

// import "fmt"

// func main() {

// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Printf("%d sayısı çiftdir.", i)
// 		} else {
// 			fmt.Printf("%d sayısı tekdir.", i)

// 		}
// 	}

// }

// 1-50 arasındaki asal sayıları yazan program

package main

import "fmt"

func main() {

	var x, y int

	for x = 2; x < 50; x++ {
		for y = 2; y < (x / y); y++ {
			if x%y == 0 {
				break
			}
		}

		if y > (x / y) {
			fmt.Printf(" Asal sayı : %d \n", x)
		}
	}
}
