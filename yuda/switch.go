package main

import "fmt"

func main() {

	a := 6

	if a >= 0 && a%3 == 0 {
		fmt.Println("nilai a habis dibagi 3")
	} else {
		fmt.Println("nilai b tidak habis dibagi 3")
	}

	switch {
	case a >= 0 && a%3 == 0:
		fmt.Println("nilai a memakai sitch habis dibagi 3")
	default:
		fmt.Println("nilai b memakai switch tidak habis dibagi 3")
	}

}
