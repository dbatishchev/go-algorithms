package sorting

import "math/rand"

func QuickSort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	pivot := s[rand.Intn(len(s))]

	l := []int{}
	m := []int{}
	r := []int{}

	for _, v := range s {
		if (v < pivot) {
			l = append(l, v)
		}
		if (v == pivot) {
			m = append(m, v)
		}
		if (v > pivot) {
			r = append(r, v)
		}
	}

	l = QuickSort(l)
	r = QuickSort(r)

	result := append(l, m...)
	result = append(result, r...)

	return result
}