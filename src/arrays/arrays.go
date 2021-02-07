package arrays

// Soma will return the sum of array numbers
func Soma(numbers []int) (soma int) {
	for _, number := range numbers {
		soma += number
	}
	return soma
}
