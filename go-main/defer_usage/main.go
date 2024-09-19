// package main

// import "fmt"

// func main() {

// 	x := 10
// 	y := 20

// 	defer fmt.Println("defer x: ", x)
// 	defer fmt.Println("defer y: ", y)

// 	x = 100
// 	y = 200

// 	fmt.Println("x : ", x)
// 	fmt.Println("y : ", y)

// }

package main

import "fmt"

func main() {
	var condition = true

	defer cleanUp()
	if condition {
		panic("An error occurred")
	}
}

func cleanUp() {
	fmt.Println("Cleanup worked...")
}
