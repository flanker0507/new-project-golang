package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d", "e", "f"}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	fmt.Println("==============")

	for index, value := range a {
		fmt.Println("index menggunakan range: ", index)
		fmt.Println("value menggunakan range: ", value)
	}
	fmt.Println("==============")

	for index := range a {
		fmt.Println("value emngguanakn range2", a[index])
	}

	fmt.Println("==============")

	b := map[string]int{"mobil": 1000, "motor": 2000, "helikopter": 3000}

	for i, v := range b {
		fmt.Println("index", i)
		fmt.Println("value", v)

	}
}
