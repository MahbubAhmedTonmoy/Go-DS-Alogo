package sort

import (
	"GoDsAlgo/constraints"
)

func getNextGap(gap int) int {
	gap = (gap * 10) / 13 //1.3
	if gap < 1 {
		return 1
	}
	return gap
}
func CombSort[T constraints.Ordered](arr []T) []T {
	n := len(arr)
	gap := n
	swapped := true

	for gap != 1 || swapped {
		gap = getNextGap(gap)
		swapped = false
		for i := 0; i < n-gap; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				swapped = true
			}
		}
	}
	return arr
}

// func main() {
// 	arr := []int{2, 3, 4, 1, 8, 9}
// 	fmt.Println(CombSort(arr))
// }
