package main

import "fmt"

func main() {
	currecny := map[string]int{
		"USD": 15000,
		"GPB": 10000,
		"IDR": 20000,
	}

	fmt.Println(currecny["USD"])
	delete(currecny, "USD")
	fmt.Println(currecny)

	b := make(map[string]int)
	b["Semangka"] = 1000
	b["Durian"] = 2000
	b["Anggur"] = 5000

	fmt.Println(b)
	fmt.Println("Harga buah semangka: ", b["Semangka"])

	value, isB := b["Durian"]
	fmt.Println("Value: ", value)
	fmt.Println("Check Durian", isB)

}
