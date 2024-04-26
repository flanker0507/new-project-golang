package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SplitWords("Yuda Wira Pratama"))

	var buah = [5]string{"apple", "grape", "banana", "melon", "watermelon"}
	fmt.Println(buah[1])

	fmt.Println("== Perulangan Elemen Array Menggunakan Keyword for==")

	for i := 0; i < len(buah); i++ {
		fmt.Printf("elemen %d : %s\n", i, buah[i])
	}

	fmt.Println("==Perulangan Elemen Array Menggunakan Keyword for - range==")

	for i, buahh := range buah {
		fmt.Printf("elemen %d : %s\n", i, buahh)
	}

	fmt.Println("==Penggunaan Variabel Underscore _ Dalam for - range==")

	for _, buahh := range buah {
		fmt.Printf("nama buah : %s\n", buahh)
	}

}

func SplitWords(input string) []string {
	return strings.Fields(input)
}
