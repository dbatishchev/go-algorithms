package utils

func Swap(s []int, i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}