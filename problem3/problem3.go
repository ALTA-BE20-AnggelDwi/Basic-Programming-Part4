package main

import "fmt"

func JoinArrayRemoveDuplicate(arrayA, arrayB []string) []string {
	// your code here
	// Membuat peta untuk melacak elemen yang sudah ditemui
	seen := make(map[string]bool)

	// Membuat slice baru untuk menyimpan hasil gabungan tanpa duplikat
	result := make([]string, 0)

	// Menambahkan elemen dari arrayA ke hasil
	for _, element := range arrayA {
		if !seen[element] {
			result = append(result, element)
			seen[element] = true
		}
	}

	// Menambahkan elemen dari arrayB ke hasil
	for _, element := range arrayB {
		if !seen[element] {
			result = append(result, element)
			seen[element] = true
		}
	}

	return result

}

func main() {
	// Test cases
	fmt.Println(JoinArrayRemoveDuplicate([]string{"apel", "anggur"}, []string{"lemon", "leci", "nanas"}))
	// ["apel", "anggur", "lemon", "leci", "nanas"]

	fmt.Println(JoinArrayRemoveDuplicate([]string{"samsung", "apple"}, []string{"apple", "sony", "xiaomi"}))
	// ["samsung", "apple", "sony", "xiaomi"]

	fmt.Println(JoinArrayRemoveDuplicate([]string{"football", "basketball"}, []string{"basketball", "football"}))
	// [football basketball]
}
