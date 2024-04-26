package main

import "fmt"

const (
	dollarToRupiah    = 15000
	euroToRupiah      = 17000
	gpbToRupiah       = 20000
	jpyToRupiah       = 104
	maxDollarExcahnge = 1000
)

func main() {
	saldoDollar := 12
	//fmt.Printf("masukkan jumlah uang: %d\n", saldoDollar)
	//fmt.Scan(&saldoDollar)

	fmt.Println("saldo dollar awal : ", saldoDollar)

	saldoRupiah := exchangeDollarToRupiah(float64(saldoDollar))

	fmt.Println("Saldo rupiah : ", saldoRupiah)

}

func exchangeDollarToRupiah(dollar float64) (rupiah float64) {
	if dollar > maxDollarExcahnge {
		fmt.Println("Maaf tidak bisa menukar lebih dari 100")
	}
	rupiah = dollar * dollarToRupiah
	rupiah = dollar * euroToRupiah
	rupiah = dollar * gpbToRupiah
	rupiah = dollar * jpyToRupiah

	return
}
