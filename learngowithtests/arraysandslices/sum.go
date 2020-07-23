package arraysandslices

func Sum(numbers []int) int {
	sum := 0;
	// range lets you iterate over an array. Every time it is called it returns two values, the index and the value.
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Go can let you write variadic functions that can take a variable number of arguments.
// Variadic Functions: https://gobyexample.com/variadic-functions
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}


/**
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
 */


// requirement:
// change SumAll to SumAllTails, where it now calculates the totals of the "tails" of each slice.
// The tail of a collection is all the items apart from the first one (the "head")


func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// Slices can be sliced! The syntax is slice[low:high] If you omit the value on one of the sides of the :
			// it captures everything to the side of it. In our case, we are saying "take from 1 to the end" with numbers[1:].
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return sums
}