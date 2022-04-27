package piscine

import (
	"ft"
)

func PrintComb() {
	isFirst := true
	nums := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for first := 0; first <= 9; first++ {
		for second := 0; second <= 9; second++ {
			for third := 0; third <= 9; third++ {
				if first < second && second < third {
					if !isFirst {
						ft.PrintRune(',')
						ft.PrintRune(' ')
					}
					ft.PrintRune(nums[first])
					ft.PrintRune(nums[second])
					ft.PrintRune(nums[third])
					isFirst = false
				}
			}
		}
	}
	ft.PrintRune('\n')
}
