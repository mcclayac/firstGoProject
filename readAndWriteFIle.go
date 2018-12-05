package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Read and write file")

	b, err := ioutil.ReadFile("inputfile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println()
	fmt.Println(string(b)) // Byte array converted to a string

	//write
	stringContent := string(b) + "\nMore New Stuff"

	fmt.Println("Write to file", stringContent)
	err = ioutil.WriteFile("outputFile.txt", []byte(stringContent), 0644)
	if err != nil {
		panic(err)
	}

}
