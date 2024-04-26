package main

import "fmt"

func main() {
	name := "dilla"

	if name == "Yuda" {
		fmt.Println("Ya Benar")
	} else {
		fmt.Println("Nama tidak sesuai")
	}

	fmt.Println("=====================")

	nilai1 := 10
	nilai2 := 20

	result := nilai1 + nilai2

	if result == 20 {
		fmt.Println("hasil sama dengan 20")
	} else if result > 20 {
		fmt.Println("hasil lebih besar dari 20")
	} else {
		fmt.Println("hasil kurang dari 20")
	}

	fmt.Println("==Variabel Temporary==")

	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	fmt.Println("==Pemanfaatan case Untuk Banyak Kondisi==")

	var poin = 6

	switch poin {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	fmt.Println("==Kurung Kurawal Pada Keyword case & default==")

	var pointt = 6

	switch pointt {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
		fmt.Println("you can be better!")
	}

	fmt.Println("==Switch Dengan Gaya if - else==")

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
		fmt.Println("you can be better!")
	}

	fmt.Println("==Seleksi Kondisi Bersarang==")

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep tyring")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder !")
			}
		}
	}

}
