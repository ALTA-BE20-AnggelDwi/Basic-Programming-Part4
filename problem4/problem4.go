package main

import (
	"fmt"
	"strconv"
)

func MunculSekali(angka string) []int {
	// your code here
	// Membuat peta untuk melacak jumlah kemunculan setiap karakter
	countMap := make(map[rune]int)

	// Menghitung kemunculan setiap karakter
	for _, char := range angka {
		countMap[char]++
	}

	// Membuat slice untuk menyimpan karakter yang hanya muncul sekali
	uniqueChars := make([]int, 0)

	// Memeriksa setiap karakter dan menambahkannya ke slice jika muncul sekali
	for _, char := range angka {
		if countMap[char] == 1 {
			// Mengonversi karakter ke bilangan bulat menggunakan strconv
			num, err := strconv.Atoi(string(char))
			if err == nil {
				uniqueChars = append(uniqueChars, num)
			}
		}
	}

	return uniqueChars

}

func main() {
	fmt.Println(MunculSekali("1234123"))    // [4]
	fmt.Println(MunculSekali("76523752"))   // [6, 3]
	fmt.Println(MunculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(MunculSekali("1122334455")) // []
	fmt.Println(MunculSekali("0872504"))    // [8 7 2 5 4]
}
