package main

import (
	"fmt"
	"os"
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

	fmt.Printf("%T", f) // print the variable type
	fmt.Printf("%T", x) // print the variable type

	fmt.Printf("%v", f) // print the variable type
	fmt.Printf("%v", x) // print the variable type

	//-------------------------------------------------
	fmt.Println("")
	carLotA := 10
	carLotB := 21 // 20
	//carlotC := 25

	carlotTotal := carLotA + carLotB
	fmt.Println("carlotTotal = ", carlotTotal)

	carlotTotal = carlotTotal + 10
	carlotTotal += 10 // Add 10
	fmt.Println("carlotTotal = ", carlotTotal)

	carlotTotal = carlotTotal * 10
	carlotTotal *= 10 // multiply by 10
	fmt.Println("carlotTotal = ", carlotTotal)

	carlotTotal = carLotB / carLotA
	//carlotTotal *= 10      // multiply by 10
	fmt.Println("carlotTotal = ", carlotTotal)

	carlotTotal = carLotB % carLotA // remainder
	//carlotTotal *= 10      // multiply by 10
	fmt.Println("carlotTotal = ", carlotTotal)

	//-------------------------------------------------
	fmt.Println("Hello World!")
	d := dealerShip{name: "A1 Auto", city: "Los Angeles"}
	d.city = "Phoenix"
	fmt.Println("city = " + d.city)
	fmt.Println("name =" + d.name)

	var d2 dealerShip
	d2 = dealerShip{name: "Discount Auto", city: "Houston"}
	fmt.Println(d2)

	d3 := dealerShip{}
	fmt.Println("d3.name=" + d3.name)

	//-------------------------------------------------
	// ------------------------------------------------
	df1 := dealerShip{"A1 Auto", "Los Angeles"}
	dealerName := creareDealerFullName(df1.name, df1.city)
	fmt.Println(dealerName)

	// ------------------------------------------------

	sold, name := calculateInvUtil(100, 175, df1)
	fmt.Println(name, sold)

	// ------------------------------------------------

	dg1 := "A1 Auto"
	dg2 := "Discount Auto"
	dg3 := "Riverside Automart"

	printDealers(dg1, dg2, dg3)
	fmt.Println("")
	dealers := []string{dg1, dg2, dg3} /// Array Declaration
	printDealers(dealers...)

	// ------------------------------------------------
	// Defer  similar to try catch finally

	file := createFile("dealers.txt")
	defer closeFile(file)
	writeToFile(file, "A1 Auto")
	writeToFile(file, "Tony McClay")

}

type dealerShip struct {
	name string
	city string
}

func creareDealerFullName(s1 string, s2 string) string {
	return s1 + " of " + s2
}

func calculateInvUtil(remaining float32, start float32, dealer dealerShip) (percentSold float32, dealerName string) {

	dealerName = dealer.name + " sold "
	percentSold = 1 - remaining/start
	return
}

func printDealers(dealers ...string) {
	for _, dealerName := range dealers { // for loop with throwaway index
		fmt.Println(dealerName)
	}
}

func createFile(path string) *os.File {
	fmt.Println("crearing file")
	file, err := os.Create(path)
	if err != nil { // if error not equal to nil ... panic (terminate)
		panic(err)
	}
	return file
}

func writeToFile(file *os.File, dealerName string) {
	fmt.Println("Writng to file")
	fmt.Fprintln(file, dealerName)
}

func closeFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()

}
