package function

type transformer func(int) int

func Function() {
	numbers := []int{1, 2, 3, 4, 5}
	squaredNumbers := transform(numbers, func(n int) int {
		return n * n
	})
	for _, n := range squaredNumbers {
		println(n)
	}
}

func transform(numbers []int, transformer func(int) int) []int {
	var result []int
	for _, n := range numbers {
		result = append(result, transformer(n))
	}
	return result
}

func getTransformer() transformer {
	return func(n int) int {
		return n * n
	}
}

// closure
func createTransformer(factor int) transformer {
	return func(n int) int {
		return n * factor
	}
}

// recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// variadic function, ... creates a slice of the multiple arguments
func sumup(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
