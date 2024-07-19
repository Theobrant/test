package main

import "fmt"

func RoundToNext5(n int) int {
	if n <= 5 {
		return 5
	} else if n > 5 && n <= 15 {
		return 15
	}
	return 1
}

func main() {
	fmt.Println("ky")
}
