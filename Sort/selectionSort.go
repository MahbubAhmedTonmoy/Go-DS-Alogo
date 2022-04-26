package sort

import (
	"GoDsAlgo/constraints"
)

func SelectionSort[T constraints.Ordered](arr []T) []T {

	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

// func main() {
// 	arr := []int{2, 3, -4, 4, 1, 8, 9, 0}
// 	fmt.Println(SelectionSort(arr))
// }
