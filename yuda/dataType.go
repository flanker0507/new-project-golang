package main

import "fmt"

func main() {
	// Int
	age := 23

	// string
	name := "Yuda Wira Pratama"

	//Float
	suhu := 35.6

	//Byte
	var BinaryData byte = 0b11111111

	//Boolean
	isMentor := true

	//nil

	var noValue interface{}

	fmt.Println("Usia", age)
	fmt.Println("Name: ", name)
	fmt.Println("Suhu: ", suhu)
	fmt.Println("Data Binar: ", BinaryData)
	fmt.Println("Adalah Mentor: ", isMentor)
	fmt.Println("Nill: ", noValue)

}
