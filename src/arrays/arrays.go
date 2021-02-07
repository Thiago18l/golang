package arrays

// Soma will return the sum of array numbers
func Soma(numbers [5]int) (soma int) {
	for i := 0; i < 5; i++ {
		soma += numbers[i]
	}
	return soma
}
