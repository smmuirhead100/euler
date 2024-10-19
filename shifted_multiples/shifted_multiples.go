package shifted_multiples

import (
	"fmt"
	"strconv"
)

func S(num int) int {
	// Handle single digit numbers
	if num < 10 {
		return num
	}

	numStr := strconv.Itoa(num)
	shifted := numStr[1:] + numStr[:1]

	// Now convert back to number
	result, _ := strconv.Atoi(shifted)

	return result
}

func N(num float64) int {
	// idk how to do this one
	for i := 1; i < 999999999999999999; i++ {
		s := S(i)
		product := int(num) * i

		if s == product {
			fmt.Println(i)
			return i
		}
	}
	return 0
}

func AreCoprime(u int, v int) bool {
	// Returns True if the two numbers being passed are coprime
	// That is, the only common divisor between them is 1.
	if u == 0 || v == 0 {
		return false
	}

	for i := 2; i <= min(u, v); i++ {
		if u%i == 0 && v%i == 0 {
			return false
		}
	}

	return true
}

func min(a, b int) int {
	// Helper to get min from two integers
	if a < b {
		return a
	}
	return b
}

func T(m int) int {
	sum := 0

	for u := 0; u < m; u++ {
		for v := 0; v < m; v++ {
			if AreCoprime(u, v) {
				uCubed := u * u * u
				vCubed := v * v * v
				fraction := float64(uCubed / vCubed)
				fmt.Println("Getting stuff for: ", fraction)
				sum += N(fraction)
			}
		}
	}

	return sum % 1000000007
}
