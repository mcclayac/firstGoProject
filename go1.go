package main

import (
	"fmt"
	"time"
)

//var carTypeDist1 string = "Toyota"
//var carTypeDist2 = "Ford"

//var carTypeDist1, carTypeDist2  string = "Toyota", "Ford"

//var carTypeDist1, carTypeDist2 = "Toyota", "Ford"

var (
	carTypeDist1 = "Toyota"
	carTypeDist2 = "Ford"
)

//const toyotaHeadQuarters  = "Toyota Motor Corporation"
//const fordHeadQuarters  = "Ford Motor Company "

//const (
//	toyotaHeadQuarters  = "Toyota Motor Corporation"
//	fordHeadQuarters  = "Ford Motor Company "
//)

const toyotaHeadQuarters, fordHeadQuarters = "Toyota Motor Corporation", "Ford Motor Company "

//var newHQ string = "NewHQ"

func main() {
	fmt.Println("Hello Tony, these nuts, 世界")
	fmt.Println("This is the Golong")
	fmt.Printf("2 + 2 = %v", 2+2)
	fmt.Println(time.Now())

	//-------------------------------------------------

	fmt.Println(carTypeDist1)
	fmt.Println(carTypeDist2)

	//-------------------------------------------------
	//toyotaHeadQuarters = "Ford HQ"   error cannot assign const

	newHQ := "Tony's Home" // In-line assignment
	const myConst = "No Value"

	fmt.Println(toyotaHeadQuarters)
	fmt.Println(fordHeadQuarters)
	fmt.Println(newHQ, myConst)
	//fmt.Println(myConst)
}
