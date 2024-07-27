package main

import "fmt"

func main() {

	fruit := []string{"Apple", "Banana", "Grape"}
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(fruit)
	fmt.Println(number)

	fruit = append(fruit, "Melon")
	fmt.Println(fruit)

	for i := 0; i < len(number); i++ {
		if number[i]%2 == 0 {
			fmt.Println("genap", number[i])
		} else {
			fmt.Println("ganjil", number[i])
		}
	}

	fmt.Println("==Bilangan Ganjil dan Genap==")

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
			fmt.Println("Bilangan Genap", i)
		} else {
			fmt.Println("Bilangan Ganjil", i)

		}
	}

}
