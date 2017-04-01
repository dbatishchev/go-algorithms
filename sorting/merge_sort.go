package sorting

func Merge(l, r[]int) []int {
	m := []int{}

	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(m, r...)
		}
		if len(r) == 0 {
			return append(m, l...)
		}
		if l[0] <= r[0] {
			m = append(m, l[0])
			l = l[1:]
		} else {
			m = append(m, r[0])
			r = r[1:]
		}
	}

	return m
}

func MergeSort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	c := len(s) / 2

	l := MergeSort(s[c:])
	r := MergeSort(s[:c])

	return Merge(l, r)
}