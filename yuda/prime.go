package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&number)

	if number <= 1 {
		fmt.Println("Angka tersebut bukan bilangan prima.")
	} else {
		isPrime := true
		var factors []int

		for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
			if number%i == 0 {
				isPrime = false
				factors = append(factors, i)
				if i != number/i {
					factors = append(factors, number/i)
				}
			}
		}

		if isPrime {
			fmt.Println("Angka tersebut adalah bilangan prima.")
		} else {
			fmt.Printf("Angka tersebut bukan bilangan prima. Faktor-faktornya adalah: %v\n", factors)
		}
	}
}
