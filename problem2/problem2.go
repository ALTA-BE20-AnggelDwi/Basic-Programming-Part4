package main

import "fmt"

func Pow(x, n int) int {
	// your code here
	result := 1

	for i := 0; i < int(n); i++ {
		result *= x
	}

	return result

}

func main() {
	fmt.Println(Pow(2, 3))  // 8
	fmt.Println(Pow(7, 2))  // 49
	fmt.Println(Pow(10, 5)) // 100000
	fmt.Println(Pow(17, 6)) // 24137569
	fmt.Println(Pow(5, 3))  // 125
}
