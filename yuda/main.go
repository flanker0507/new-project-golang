package main

import "fmt"

func main() {
	//fmt.Println("Hello, Go")
	//
	//fmt.Println("=================")
	//
	//name := "Yuda Wira Pratama"
	//age := 23
	//isMentor := true
	//fmt.Println("Name: ", name)
	//fmt.Println("Umur: ", age)
	//fmt.Println("Is Mentor: ", isMentor)
	//
	//fmt.Println("=================")
	//
	//var radius float64 = 3.0
	//var area float64
	//
	//const pi = 3.14
	//
	//area = pi * (radius * radius)
	//
	//fmt.Println(area)

	for i := 1; i <= 100; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println(i, "FizzBUzz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		}
	}

}
