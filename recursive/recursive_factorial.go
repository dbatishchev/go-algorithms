package recursive

func Factorial(i int) int {
	if i == 1 {
		return i
	}
	return i * Factorial(i - 1)
}