package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "murat"
	names[1] = "mustafa"
	names[2] = "ibrahim"
	fmt.Println(names)
	names[2] = "ahmet"
	fmt.Println(names)

	var scores = []int{1, 2, 3, 4, 5}
	fmt.Println(scores)
	scores = append(scores, 6)
	fmt.Println(scores)

}
