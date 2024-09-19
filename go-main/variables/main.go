package main

import (
	"fmt"
)

func main() {

	// var productName string
	// var quantity int

	// productName = "Golang dersleri"
	// quantity = 5

	// fmt.Println(productName, reflect.TypeOf(productName))
	// fmt.Println(quantity, reflect.TypeOf(quantity))

	// var result string = "sonuc"
	// fmt.Println(result)

	// var isActive bool = false
	// fmt.Println(isActive)

	// var myFloat float64 = 3.14
	// fmt.Println(myFloat)

	// myString := "My String"
	// fmt.Println(myString)

	//? GİBİ DEĞİŞKENLER TANIMLANABİLİR

	var productName string
	var quantity int
	var discount float32
	var isInSelect bool

	productName = "Golang dersleri"
	quantity = 5
	discount = 0.37
	isInSelect = true

	fmt.Println("Product Name: ", productName, "Quantity : ", quantity, "Discount : ", discount, "İsInSelect : ", isInSelect)
	fmt.Printf("Product Name: %s Quantity : %d Discount: %f isInSelect : %t", productName, quantity, discount, isInSelect)

}
