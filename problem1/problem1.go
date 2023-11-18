package main

import (
	"fmt"
	"math"
)

func PrimeNumber(number int) bool {
	// your code here
	if number <= 1 {
		return false
	} else if number == 2 {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(PrimeNumber(1000000007))  // true
	fmt.Println(PrimeNumber(1500450271))  // true
	fmt.Println(PrimeNumber(1000000000))  // false
	fmt.Println(PrimeNumber(10000000019)) // true
	fmt.Println(PrimeNumber(10000000033)) // true
}
