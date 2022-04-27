package piscine

func pow10(n int) int {
	ret := 1
	for i := 0; i < n; i++ {
		ret *= 10
	}
	return ret
}

func PrintCombN(n int) {
	max := pow10(n)
	for i := 0; i < max; i++ {

	}
}
