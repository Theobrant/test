package main

func RoundToNext5(n int) int {
	if n <= 5 {
		return 5
	} else if n > 5 && n <= 15 {
		return 15
	}
	return 0
}

func main() {
}
