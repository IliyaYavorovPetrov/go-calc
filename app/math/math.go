package math

func Sum(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}

	return result
}

func Sub(numbers []int) int {
	result := 0
	counter := 0

	for _, num := range numbers {
		if counter == 0 {
			result = num
		} else {
			result -= num
		}

		counter++
	}

	return result
}