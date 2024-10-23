package all_euler

import (
	"math/big"
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
	// Pandigital: Each digit can only exist once in the number
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

// ----------------
// Counting Sundays

var MONTHS_TO_DAYS = map[string]int{
	"January":   31,
	"February":  28,
	"March":     31,
	"April":     30,
	"May":       31,
	"June":      30,
	"July":      31,
	"August":    31,
	"September": 30,
	"October":   31,
	"November":  30,
	"December":  31,
}

var MONTHS_ORDER = []string{
	"January", "February", "March", "April", "May", "June",
	"July", "August", "September", "October", "November", "December",
}

func IsLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 {
		return true
	} else if year%400 == 0 {
		return true
	}
	return false
}

func CountingSundaysOnFirstSince(startMonth string, startYear int, endMonth string, endYear int) int {
	offset := 1
	monthIndex := 0
	year := 1900

	// Determine offset to start at
	for {
		month := MONTHS_ORDER[monthIndex]
		if startYear == year && startMonth == month {
			break
		}

		leftover := MONTHS_TO_DAYS[month] % 7
		offset += leftover
		offset = offset % 7

		monthIndex++
		if monthIndex == 12 {
			year++
			monthIndex = 0
		}

	}

	sundaysOnFirstOfMonth := 0
	if offset == 0 {
		sundaysOnFirstOfMonth++
	}
	for {
		month := MONTHS_ORDER[monthIndex]
		days := MONTHS_TO_DAYS[month]

		if endYear == year && startMonth == month { // TODO: This is not accurate.
			break
		}

		// Account for leap years
		if month == "February" {
			if IsLeapYear(year) {
				days = 29
			}
		}

		leftover := days % 7
		offset += leftover
		offset = offset % 7

		// This is key - means that we landed on a Sunday!
		if offset == 0 {
			sundaysOnFirstOfMonth += 1
		}

		if month == "December" {
			year++
			monthIndex = 0
		} else {
			monthIndex++
		}
	}

	return sundaysOnFirstOfMonth
}

func DistinctPowers(min, max int) int {
	distinctPowers := make(map[string]bool)
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			base := big.NewInt(int64(a))
			exp := big.NewInt(int64(b))
			result := new(big.Int).Exp(base, exp, nil)

			// Convert to string to use as map key
			resultStr := result.String()
			distinctPowers[resultStr] = true
		}
	}
	return len(distinctPowers)
}
