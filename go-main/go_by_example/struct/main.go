package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var person1 Person
	var person2 Person

	// person1
	person1.name = "İbrahim"
	name := person1.name
	person1.age = 23
	person1.job = "Yazılım Mühendisi"
	person1.salary = 40000

	// person2
	person2.name = "Fatih"
	person2.age = 27
	person2.job = "Tasarımcı"
	person2.salary = 50000

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", person1.age)
	fmt.Println("Job: ", person1.job)
	fmt.Println("Salary: ", person1.salary)
	fmt.Println("-------------------")
	fmt.Println("Name: ", person2.name)
	fmt.Println("Age: ", person2.age)
	fmt.Println("Job: ", person2.job)
	fmt.Println("Salary: ", person2.salary)

}
