package piscine

import (
	"ft"
)

func printNum(num int) {
	nums := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	if num >= 10 {
		ft.PrintRune(nums[num/10])
		ft.PrintRune(nums[num%10])
	} else {
		ft.PrintRune(nums[0])
		ft.PrintRune(nums[num])
	}
}

func PrintComb2() {
	isFirst := true
	for first := 0; first <= 99; first++ {
		for second := 0; second <= 99; second++ {
			if first < second {
				if !isFirst {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
				printNum(first)
				ft.PrintRune(' ')
				printNum(second)
				isFirst = false
			}
		}
	}
	ft.PrintRune('\n')
}
