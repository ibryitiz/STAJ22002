package main

import "fmt"

func main() {
	var customer1 = Customer{id: 1, name: "İbrahim", age: 23, address: Address{city: "İstanbul", district: "Maltepe"}}

	customer1.ToString()
	customer1.changeName("Ahmet")
	customer1.ToString()
}

//! burada yine name değişmez referansdan(pointer) değiştirmek lazım
// func (customer Customer) changeName(newName string) {
// 	customer.name = newName
// }

//* structa ait fonksiyon ve referans(pointer) üzerinden değiştirme.
func (customer *Customer) changeName(newName string) {
	customer.name = newName
}

func (customer *Customer) ToString() {
	fmt.Printf("Name : %s , Age : %d\n", customer.name, customer.age)
}

type Customer struct {
	id      int64
	name    string
	age     int
	address Address
}

type Address struct {
	city     string
	district string
}
