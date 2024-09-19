package main

import "fmt"

func main() {
	x := 5
	fmt.Println(x) // 5
	double(&x)     // func ile 10 yaptık
	fmt.Println(x) // 10 oldu 5 olarak kalmadı
}

func double(num *int) {
	*num = *num * 2
	fmt.Println(num)  // adres yazar
	fmt.Println(*num) // deger yazar

}
