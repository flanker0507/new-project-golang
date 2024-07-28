package main

import "fmt"

func main() {
	fmt.Println("== hitung rata-rata ==")
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var total float64
	count := len(scores)
	var goodScores []int

	for _, score := range scores {
		total = total + float64(score)
	}

	avarage := total / float64(count)
	fmt.Printf("Rata-rata: %.2f\n", avarage)

	fmt.Println("========================")

	for _, score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}
	fmt.Println(goodScores)

}
