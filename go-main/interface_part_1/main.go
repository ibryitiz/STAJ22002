package main

import "fmt"

type IShippable interface {
	shippingCost() int
}

//! ÖNEMLİ
// Bu interface referanslı fonksiyon sayesinde
// bir tane tek fonksiyon olucak ve herhangi bir Book veya yeni bir struct bile
// gelirse shippingCost() fonksiyonunu kullanan buna erişebilecek.

func calculateTotalShippingCost(products []IShippable) int {
	total := 0

	for _, product := range products {
		total += product.shippingCost()
	}
	return total
}

func main() {

	// book1 := &Book{desi: 10}
	// book2 := &Book{desi: 20}

	// fmt.Println(book1.shippingCost())
	// fmt.Println(book2.shippingCost())

	// books := []Book{*book1, *book2}
	// fmt.Println(calculateTotalShippingCostOfBooks(books))

	// fmt.Println("-----------")

	// electronic1 := &Electronic{desi: 20}
	// electronic2 := &Electronic{desi: 30}

	// fmt.Println(electronic1.shippingCost())

	// electronics := []Electronic{*electronic1, *electronic2}
	// fmt.Println(calculateTotalShippingCostOfElectronics(electronics))

	//!İNTERFACE

	var product IShippable

	product = &Book{desi: 10}
	fmt.Println(product.shippingCost())
	//* Artık her iki structa da tek bir değişken ile ulaşıyoruz
	product = &Electronic{desi: 20}
	fmt.Println(product.shippingCost())

	fmt.Println("----------------------")

	var products []IShippable = []IShippable{
		&Book{desi: 10},
		&Electronic{desi: 20},
		&Book{desi: 30},
		&Electronic{desi: 40},
		// Flower eklendi ve interface ile ilgili kod değişikliği yapılmadan totali hesapladık.
		&Flower{desi: 50},
	}

	fmt.Println(calculateTotalShippingCost(products))
}

type Book struct {
	desi int
}

func (book *Book) shippingCost() int {
	return 5 + book.desi*2
}

//! Gerek yok
// func calculateTotalShippingCostOfBooks(books []Book) int {
// 	total := 0

// 	for _, book := range books {
// 		total += book.shippingCost()
// 	}
// 	return total
// }

type Electronic struct {
	desi int
}

func (electronic *Electronic) shippingCost() int {
	return 10 + electronic.desi*3
}

//! Gerek yok
// func calculateTotalShippingCostOfElectronics(electronics []Electronic) int {
// 	total := 0

// 	for _, electronic := range electronics {
// 		total += electronic.shippingCost()
// 	}
// 	return total
// }

type Flower struct {
	desi int
}

func (flower *Flower) shippingCost() int {
	return 15 + flower.desi*4
}
