package sort

import (
	"GoDsAlgo/constraints"
)

func Partition[T constraints.Ordered](arr []T, low, high int) int {
	index := low - 1
	pivotElement := arr[high]
	for i := low; i < high; i++ {
		if arr[i] <= pivotElement {
			index++
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}

func QuickSortRange[T constraints.Ordered](arr []T, low, high int) {
	if len(arr) <= 1 {
		return
	}

	if low < high {
		pivot := Partition(arr, low, high)

		QuickSortRange(arr, low, pivot-1)
		QuickSortRange(arr, pivot+1, high)
	}
}

func QuickSort[T constraints.Ordered](arr []T) []T {
	QuickSortRange(arr, 0, len(arr)-1)
	return arr
}

// func main() {
// 	arr := []int{2, 3, -4, 4, 1, 8, 9, 0}
// 	fmt.Println(QuickSort(arr))
// }
