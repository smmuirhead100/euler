package all_euler

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
