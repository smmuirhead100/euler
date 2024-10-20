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

func SumSquareDifference(num int) int {
	sumSqSum := 0
	sum := 0
	for i := 0; i <= num; i++ {
		sumSqSum += i * i
		sum += i
	}
	sumSq := sum * sum

	return sumSq - sumSqSum
}

func GetDivisors(num int) []int {
	// Return list of all valid divisors for given int
	nums := []int{}

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			nums = append(nums, i)
		}
	}
	return nums
}

func getTriangleNumber(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

func HighlyDivisibleTriangularNumber(num int) int {
	b := 1
	for {
		sum := getTriangleNumber(b)
		divisors := GetDivisors(sum)
		if len(divisors) > num {
			return sum
		}
		b++
	}
}

func PanDigitalMultiples(target int) int {
	// 38
	closestToTarget := 0
	for i := 0; i < target; i++ {
		currentSum := 0
		for n := 0; n < 10; n++ {
			product := i * n
			currentSum += product

			currentDifference := target - closestToTarget
			contendingDifference := target - currentSum

			if contendingDifference < 0 {
				break
			} else if currentDifference > contendingDifference {
				closestToTarget = currentSum
			}
		}
	}
	return closestToTarget
}
