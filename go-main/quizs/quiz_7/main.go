// Soru 1
// 5 string elemandan oluşan bir array
// ve 5 int elemandan
// oluşan slice oluşturup index numaralarıyla
// birlikte gösterin

package main

import "fmt"

func main() {
	var texts = [5]string{"text0", "text1", "text2", "text3", "text4"}

	var numbers = []int{0, 1, 2, 3, 4}

	for _, text := range texts {
		for _, number := range numbers {
			fmt.Printf("%d . text : %s  \n", number, text)
		}
		fmt.Println("-----------------------")
	}

}
