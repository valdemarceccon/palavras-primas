package primos

func CheckWord(word string) (bool, error) {
	if len(word) == 0 {
		return false, nil
	}

	result := IsPrime(sum(word))

	return result, nil
}

func sum(word string) (sum int) {
	offset := int('a' - 1)
	sum = 0
	for _, letra := range word {
		sum += int(letra) - offset
	}

	return
}
