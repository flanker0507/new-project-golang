package main

import "fmt"

const (
	dollarToRupiah    = 15000
	maxDollarExcahnge = 100
)

func main() {
	saldoDollar := 10.0

	fmt.Println("saldo dollar awal : ", saldoDollar)

	saldoRupiah := exchangeDollarToRupiah(saldoDollar)

	fmt.Println("Saldo rupiah : ", saldoRupiah)

}

func exchangeDollarToRupiah(dollar float64) float64 {
	if dollar > maxDollarExcahnge {
		fmt.Println("Maaf tidak bisa menukar lebih dari 100")
	}
	rupiah := dollar * dollarToRupiah

	return rupiah
}
