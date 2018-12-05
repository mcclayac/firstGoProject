package main

import (
	cs "CarStuff"
	"fmt"
)

func main() {
	c := cs.Car{4, 6}
	fmt.Println(c)
	fmt.Println(c.GetDoors())
	t2 := Truck{4, "Full", oneTon}
	fmt.Println(t2)

	fmt.Println(t2.GetDoors())

}

type Truck struct {
	NumberOfDoors int
	BedSize       string
	WeightByTon   weight
}

type weight string

const (
	oneTon weight = "One Ton"
	twoTon weight = "Two Ton"
)

func (t *Truck) GetDoors() int {
	return t.NumberOfDoors
}
