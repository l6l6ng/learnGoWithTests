package sum

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAllTails(numberToSum ...[]int) (sums []int) {
	//lengthOfNumbers := len(numberToSum)
	//sums = make([]int, lengthOfNumbers)

	for _, numbers := range numberToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
