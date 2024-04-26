package main

import "fmt"

func main() {
	var a, b int
	var choice string

	fmt.Print("masukkan angka pertama: ")
	fmt.Scan(&a)
	fmt.Print("masukkan angka ke dua: ")
	fmt.Scan(&b)

	for {
		//fmt.Print("Masukkan dua angka (pisahkan dengan spasi): ")

		fmt.Printf("%d + %d = %d\n", a, b, a+b)
		fmt.Printf("%d - %d = %d\n", a, b, a-b)
		fmt.Printf("%d * %d = %d\n", a, b, a*b)
		if b != 0 {
			fmt.Printf("%d / %d = %d\n", a, b, a/b)
			fmt.Printf("%d %% %d = %d\n", a, b, a%b)
		} else {
			fmt.Println("Pembagian oleh nol tidak diizinkan.")
		}
		fmt.Print("Apakah anda ingin melakukan operasi lagi? (yes/no): ")
		fmt.Scan(&choice)
		if choice != "yes" {
			return
		}
	}

}
