package sorting

import "go-algorithms/utils"

func InsertionSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	for i := 1; i < len(s); i++ {
		for j := i; j > 0; j-- {
			if (s[j] < s[j - 1]) {
				utils.Swap(s, j, j - 1)
			}
		}
	}
	return s
}