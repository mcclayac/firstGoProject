package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
Go types

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
	 // represents a unicode code point

float32 float64

complex64 complex128

*/
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

	//-------------------------------------------------
	var mystring = "3"
	var x = 10
	var f float32 = 2.0
	fmt.Println(mystring, x, f)

	number, _ := strconv.Atoi(mystring) //Atoi returns 2 values
	fmt.Println(number)
	x = number
	fmt.Println(x)

}
