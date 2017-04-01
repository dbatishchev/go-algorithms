package sorting

func swap(s []int, i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func Sort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(s) - 1; i++ {
			if (s[i] > s[i + 1]) {
				swapped = true
				swap(s, i, i + 1)
			}
		}
	}
	return s
}