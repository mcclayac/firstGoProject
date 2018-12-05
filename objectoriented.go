package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go Object Oriented")

	// -----------------------------------------
	// Composition
	aCar := car{}
	fmt.Println(aCar)
	aCar.mpg = 25
	aCar.numberOfDoor = 2
	fmt.Println(aCar)
	// fmt.Println(aCar.getMpg())
	aCar.getMpg()

	bCar := car{vehicle{19, 4}, red} // Initialize a vehicle in a car
	fmt.Println(bCar)
	bCar.getMpg()

	// -----------------------------------------
	// Custom Types
	bCar.color = blue
	fmt.Println(bCar)

	cCar := car{}
	cCar.color = black
	cCar.mpg = 34
	cCar.numberOfDoor = 4
	fmt.Println(cCar)
	cCar.getMpg()

	// -----------------------------------------
	// Polymorphic
	aTruck := truck{vehicle{15, 2}, black}
	mans := [...]manufacturer{cCar, aTruck}
	for i, _ := range mans {
		fmt.Println("iteration=", i, "mans=", mans[i].getVin())
	}

}

type vehicle struct {
	mpg          int
	numberOfDoor int
}

type car struct {
	vehicle
	color Color
}

type truck struct {
	vehicle
	color Color
}

func (t truck) getVin() string {
	return "truck vin"
}

func (c car) getVin() string {
	return "car vin"
}

type manufacturer interface {
	getVin() string
}

func (v *vehicle) getMpg() {
	fmt.Println("This vehicle gets", v.mpg, "mpg")
}

// Custom type  struct {
type Color string

const (
	blue  Color = "blue"
	red   Color = "red"
	green Color = "green"
	black Color = "black"
)
