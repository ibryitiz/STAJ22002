package main

import "fmt"

func main() {
	var a int = 10

	fmt.Println(a)

	//! pointer tanımlama
	var p *int = &a //* burada a nın değerini p adresinde tutuyoruz

	fmt.Println(p)  //* burada tutulduğu adresi yazıyor
	fmt.Println(*p) //* burada adrese git ve o değeri getir diyoruz(10)

	// pointer değer değiştirme

	*p = 20 // burada ise o adrese git değeri 20 yap diyoruz
	// a artık 20 oldu
	fmt.Println(a)

	fmt.Println("--------------")

	pointerAnlama()

	fmt.Println("-------------")

	var w int = 10
	fmt.Println(w)                       // 10
	c := pointerNedenKullanilirNormal(w) // 22
	fmt.Println("ahahhahaha", c)
	fmt.Println(w) // 10

	fmt.Println("-----pointer --------")
	fmt.Println(w)                    // 10
	pointerNedenKullanilirPointer(&w) // 22
	fmt.Println(w)                    // 22

	//? pointer ve arraylar
	var numbers = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers)
	changeValue(numbers)
	fmt.Println(numbers)
	//! arraylar references yapıda olduğu için
	//! pointer kullanmadan slice daki değeri değiştirdik.

}

func pointerAnlama() {
	var a = 10
	var b int
	var p *int

	b = a
	p = &a

	*p = 20
	fmt.Println(a, b)
	// çıktısı  = ?
}

func pointerNedenKullanilirNormal(x int) int {
	x = x + 12
	return x
}

func pointerNedenKullanilirPointer(x *int) {
	*x = *x + 12
}

func changeValue(numbers []int) {
	numbers[0] = 1000
}
