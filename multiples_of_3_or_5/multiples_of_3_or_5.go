package multiples_of_3_or_5

func MultiplesOfThreeOrFive(num int) int {
	sum := 0
	for i := 0; i < num; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	return sum
}
