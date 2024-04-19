package main

import "fmt"

// isPalindrome checks if an integer is a palindrome
func isPalindrome(n int) bool {
	original := n
	reversed := 0

	// Handle negative numbers
	if n < 0 {
		return false
	}

	// Reverse the number
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
		//	fmt.Println("reserved: ", n)
	}

	// Check if the original number is the same as the reversed number
	return original == reversed
}

func main() {
	testNumbers := []int{121, 1221}
	for _, num := range testNumbers {
		fmt.Printf("Is %d a palindrome? %v\n", num, isPalindrome(num))
	}
}

/*
Baik, mari kita ulas logika dari fungsi pemeriksaan palindrome untuk bilangan integer dengan cara yang lebih sederhana:

### Fungsi `isPalindrome`
Fungsi ini bertujuan untuk menentukan apakah suatu bilangan integer adalah palindrome, yaitu bilangan yang membaca sama dari depan maupun dari belakang.

#### Langkah 1: Menangani Bilangan Negatif
- Bilangan negatif tidak dianggap sebagai palindrome karena tanda minus di depan tidak akan muncul di belakang saat bilangan dibalik.
- Jika bilangan yang diberikan adalah negatif, fungsi langsung mengembalikan `false`, karena tidak memenuhi kriteria palindrome.

#### Langkah 2: Membalikkan Bilangan
- **Inisialisasi variabel `reversed`:** Diawali dengan nilai 0, variabel ini akan digunakan untuk menyimpan bentuk terbalik dari bilangan asli.
- **Looping untuk membalikkan bilangan:**
- Menggunakan sisa pembagian (`n%10`) untuk mendapatkan digit terakhir dari bilangan. Contoh: Dari 123, kita mendapatkan 3.
- Mengalikan `reversed` dengan 10 dan menambahkan digit terakhir ini ke dalam `reversed`. Contoh: Jika `reversed` sudah memiliki 2 (dari langkah sebelumnya), menjadi 20 ketika dikali 10, lalu ditambah 3 menjadi 23.
- Membagi `n` dengan 10 (`n /= 10`) untuk mengurangi jumlah digitnya. Contoh: 123 menjadi 12.
- **Ulangi proses ini** hingga `n` menjadi 0, yang berarti semua digit telah diproses.

#### Langkah 3: Membandingkan Bilangan Asli dengan Bilangan Terbalik
- Setelah `n` menjadi 0, `reversed` sekarang berisi bilangan asli yang telah dibalik.
- Bandingkan bilangan asli (disimpan di `original`) dengan `reversed`. Jika keduanya sama, maka bilangan tersebut adalah palindrome, jika tidak, bukan palindrome.

### Ringkasan
Fungsi ini efektif karena:
- Menggunakan logika matematika sederhana untuk membalik bilangan.
- Membandingkan hasilnya dengan bilangan asli.
- Mengembalikan hasilnya dalam bentuk `true` atau `false` tergantung apakah bilangan tersebut palindrome atau tidak.

Ini adalah cara yang cukup efisien dan jelas untuk memeriksa palindrome pada bilangan integer dalam bahasa pemrograman Go.

*/
