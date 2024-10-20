package all_euler

import (
	"strconv"
)

func PythagoreanTriplet(num int) int {
	// Problem 9:
	// Finds the pythagorean triplet where a + b + c = num and a^2 + b^2 == c^2
	// Returns the product of the three.

	for a := 1; a < num/3; a++ { //
		for b := a; b < (num-a)/2; b++ {
			c := num - a - b
			if a*a+b*b == c*c {
				return a * b * c
			}
		}
	}
	return 0
}

func IsPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func GetPrimesByReplacing(num, startIndex, endIndex int) []int {
	// Start index (inclusive) and end index (exclusive)
	// TODO: Incomplete (Part of 55)
	primes := []int{}
	numStr := strconv.Itoa(num)

	if startIndex == endIndex {
		for i := 0; i < 10; i++ {
			replaced := numStr[:startIndex] + strconv.Itoa(i) + numStr[(startIndex+1):]
			result, _ := strconv.Atoi(replaced)
			if IsPrime(result) {
				primes = append(primes, result)
			}
		}
	} else {

	}
	return primes
}

func PrimeDigitReplacements() int {
	return 0
}

func LatticePaths(w, h int) int {
	// 2D slice
	paths := make([][]int, w+1)
	for i := range paths {
		paths[i] = make([]int, h+1)
	}

	// Edges
	for i := 0; i <= w; i++ {
		paths[i][0] = 1
	}
	for j := 0; j <= h; j++ {
		paths[0][j] = 1
	}

	// Fill the grid
	for i := 1; i <= w; i++ {
		for j := 1; j <= h; j++ {
			paths[i][j] = paths[i-1][j] + paths[i][j-1]
		}
	}
	return paths[w][h]
}
