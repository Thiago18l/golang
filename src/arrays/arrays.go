package arrays

// Soma will return the sum of array numbers
func Soma(numbers []int) (soma int) {
	for _, number := range numbers {
		soma += number
	}
	return soma
}

// SomaTudo will return a slice.
func SomaTudo(numbers ...[]int) (somas []int) {
	for _, number := range numbers {
		somas = append(somas, Soma(number))
	}
	return somas
}

// SomaTodoResto will return a slice
func SomaTodoResto(numbers ...[]int) (soma []int) {
	for _, number := range numbers {
		if len(number) == 0 {
			soma = append(soma, 0)
		} else {
			final := number[1:]
			soma = append(soma, Soma(final))
		}
	}
	return soma
}
