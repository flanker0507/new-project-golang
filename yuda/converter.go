package main

import "fmt"

const (
	dollarToRupiah    = 15000
	euroToRupiah      = 17000
	gbpToRupiah       = 20000
	jpyToRupiah       = 104
	maxDollarExchange = 1000
)

func main() {

	var saldoRupiah float64
	var currencyType, choice string

	for {
		fmt.Println("masukkan jumlah: ")
		_, err := fmt.Scan(&saldoRupiah)
		if err != nil {
			continue
		}

		if saldoRupiah > maxDollarExchange {
			fmt.Println("Maaf tidak bisa menukar lebih dari 100")
			continue
		}

		fmt.Println("Choose Currncy Type (USD, EUR, GPB, JPY)")
		_, err = fmt.Scan(&currencyType)
		if err != nil {
			continue
		}

		saldoDollar, err := exchangeDollarToRupiah(saldoRupiah, currencyType)
		if err != nil {
			continue
		}

		fmt.Printf("Balance in Rupiah: %.2f\n", saldoDollar)

		fmt.Println("do you want currenncy again? (yes/no)")
		fmt.Scan(&choice)
		if choice != "yes" {
			break
		}

	}

}

func exchangeDollarToRupiah(dollar float64, currency string) (rupiah float64, err error) {

	switch currency {
	case "USD":
		rupiah = dollar * dollarToRupiah
	case "EUR":
		rupiah = dollar * euroToRupiah
	case "GPB":
		rupiah = dollar * gbpToRupiah
	case "JPY":
		rupiah = dollar * jpyToRupiah
	default:
		fmt.Println("currency not found")
	}
	return rupiah, nil
}
