package main

import "fmt"

func main() {
	var customer1 = Customer{id: 1, name: "İbrahim", age: 23, address: Address{city: "İstanbul", district: "Maltepe"}}
	//var customer2 = Customer{id: 2, name: "Ali", age: 17}

	// fmt.Println(customer1)
	// fmt.Println(customer2)

	// fmt.Println(customer1.address.city)

	// customer1.age = 24
	// fmt.Println(customer1)

	customer1.ToString()

	fmt.Println("-----------------")

	changeName(customer1)
	customer1.ToString() //! NAME ALANI DEĞİŞMEDİ BU YÜZDEN POİNTER KULLANARAK REFERANS ÜZERİNDEN DEĞİŞTRİME YAPMALIYIZ

	changeNamePointer(&customer1)
	customer1.ToString() //* NAME ARTIK "Ahmet" OLDU.
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

//* structa ait fonksiyon tanımlama
func (customer Customer) ToString() {
	fmt.Printf("Name : %s , Age : %d\n", customer.name, customer.age)
}

func changeName(customer Customer) {
	customer.name = "Ahmet"
}
func changeNamePointer(customer *Customer) {
	customer.name = "Ahmet"
}
