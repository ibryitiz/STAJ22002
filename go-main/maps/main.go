package main

import "fmt"

func main() {

	// 1. tanım
	var names map[string]int

	names = make(map[string]int, 0)

	names["mustafa"] = 1
	names["ali"] = 2
	names["veli"] = 3

	fmt.Println(names["mustafa"])
	fmt.Println(names["melike"]) //! melike olmadığı için hata vermez 0 yazar

	// 2.tanım
	ages := map[string]int{
		"ali":  25,
		"veli": 30,
		"ayşe": 28,
	}

	fmt.Println(ages)

}
