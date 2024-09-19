package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func newPerson(name string, age int, address string) Person {
	return Person{name, age, address}
}

func main() {

	persons := make(map[string]Person)

	persons["Person 1"] = newPerson("İbrahim", 23, "Balikesir Edremit")
	persons["Person 2"] = newPerson("Ali", 16, "İstanbul Sultanahmet")
	persons["Person 3"] = newPerson("Veli", 35, "Ankara Çankaya")
	persons["Person 4"] = newPerson("Fatma", 24, "İstanbul ��işli")

	//! Benim kodum
	// var ages []int
	// for key, value := range persons {
	// 	ages := append(ages, value.age)
	// 	fmt.Printf("%s: %d\n", key, ages)
	// 	for i := 0; i < len(ages); i++ {
	// 		if ages[i] >= 30 {
	// 			fmt.Printf("%s 30 yaşından büyüktür.\n", key)

	// 		} else {
	// 			fmt.Printf("%s 30 yaşından küçüktür.\n", key)
	// 		}
	// 	}
	// }

	//! chat kodu

	var ages []int
	for key, value := range persons {
		ages = append(ages, value.age)
		fmt.Printf("%s: %d\n", key, value.age) // Yaşları doğrudan yazdırıyoruz
	}

	for i, age := range ages {
		key := fmt.Sprintf("Person %d", i+1) // Anahtarı bulmak için dizi indeksini kullanıyoruz
		if age >= 30 {
			fmt.Printf("%s 30 yaşından büyüktür.\n", key)
		} else {
			fmt.Printf("%s 30 yaşından küçüktür.\n", key)
		}
	}

}
