package sorting

import (
	"testing"
	"fmt"
)

func TestSort(t *testing.T) {
	testlist := []int{5, 6, 4, 3, 9, 8, 7, 1, 2, 4, 3, 0}
	testlist = Sort(testlist)

	for i := 0; i < len(testlist); i++ {
		fmt.Println(testlist[i])
		if (i == 0) {
			continue
		}
		if testlist[i] < testlist[i - 1] {
			t.Error("Invalid array sort")
			break
		}
	}
}