package sorting

import (
	"go-algorithms/utils"
)

func SelectionSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	var min int = 0
	for i := 0; i < len(s); i++ {
		min = i
		for j := i + 1; j < len(s); j++ {
			if (s[j] < s[min]) {
				min = j
			}
		}
		utils.Swap(s, i, min)
	}
	return s
}