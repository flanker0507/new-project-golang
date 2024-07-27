package main

import "fmt"

func main() {

	i := "Golang the best language"
	//v:= []string{"a", "i", "u", "e", "o"}

	for index, value := range i {
		//fmt.Println("index: ", index, " value: ", string(value))

		if index%2 == 0 {
			fmt.Println("index: ", index)
		}

		if string(value) == "a" {
			fmt.Println("value: ", string(value))
		} else if string(value) == "i" {
			fmt.Println("value: ", string(value))
		} else if string(value) == "u" {
			fmt.Println("value: ", string(value))
		} else if string(value) == "e" {
			fmt.Println("value: ", string(value))
		} else if string(value) == "o" {
			fmt.Println("value: ", string(value))
		}

	}

}
