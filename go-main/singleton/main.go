package main

import (
	"log"
	"onevideogo/singleton/service"
)

func main() {

	var singleton service.IdServiceSingleton

	s1 := singleton.GetService()
	buildCar(s1.Next())
	buildCar(s1.Next())

	// somewhere else in the program

	s2 := singleton.GetService()

	buildMotorBike(s2.Next())
	buildMotorBike(s2.Next())
	buildMotorBike(s1.Next())

	// s1 := service.newIdService()
	// buildCar(s1.Next())
	// buildCar(s1.Next())

	//* somewhere else in the program
	// s2 := service.newIdService()
	// buildMotorBike(s2.Next())
}

func buildCar(id int) {
	log.Print("car : ", id)
}

func buildMotorBike(id int) {
	log.Print("motorBike : ", id)
}
