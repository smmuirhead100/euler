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

func LatticePaths(w, h int, memo map[string]int) int {
	// Find number of paths you can take from top left
	// of grid to bottom right
	if w == 0 || h == 0 {
		return 1
	}

	// Memo
	key := strconv.Itoa(w) + "," + strconv.Itoa(h)
	if val, exists := memo[key]; exists {
		return val
	}

	pathsIfGoRight := LatticePaths(w-1, h, memo)
	pathsIfGoDown := LatticePaths(w, h-1, memo)

	result := pathsIfGoRight + pathsIfGoDown
	memo[key] = result

	return result
}
