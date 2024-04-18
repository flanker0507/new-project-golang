package main

import "fmt"

func main() {
	var a, b int
	var choice string

	for {
		fmt.Print("Masukkan dua angka (pisahkan dengan spasi): ")
		fmt.Scan(&a, &b)

		fmt.Printf("%d + %d = %d\n", a, b, a+b)
		fmt.Printf("%d - %d = %d\n", a, b, a-b)
		fmt.Printf("%d * %d = %d\n", a, b, a*b)
		if b != 0 {
			fmt.Printf("%d / %d = %d\n", a, b, a/b)
			fmt.Printf("%d %% %d = %d\n", a, b, a%b)
		} else {
			fmt.Println("Pembagian oleh nol tidak diizinkan.")
		}
	}
	fmt.Print("Apakah anda ingin melakukan operasi lagi? (yes/no): ")
	fmt.Scan(&choice)
	if choice != "yes" {
		return
	}

}
