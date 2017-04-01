package sorting

import "go-algorithms/utils"

func BubbleSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(s) - 1; i++ {
			if (s[i] > s[i + 1]) {
				swapped = true
				utils.Swap(s, i, i + 1)
			}
		}
	}
	return s
}