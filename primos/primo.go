package primos

func IsPrime(num int) bool {
	if num > 2 && num%2 == 0 {
		return false
	}

	for div := num - 2; div > 2; div = div - 2 {
		if num%div == 0 {
			return false
		}
	}
	return num > 1
}
