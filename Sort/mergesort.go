package sort

import (
	"GoDsAlgo/constraints"
)

func MergeSort[T constraints.Ordered](arr []T) []T {
	return Merge(arr)
}

// func MergeSort[T constraints.Ordered](items []T) []T {
// 	for step := 1; step < len(items); step += step {
// 		for i := 0; i+step < len(items); i += 2 * step {
// 			tmp := merge(items[i:i+step], items[i+step:min.Int(i+2*step, len(items))])
// 			copy(items[i:], tmp)
// 		}
// 	}
// 	return items
// }
func Merge[T constraints.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}
	var middle = len(arr) / 2
	var part1 = Merge(arr[:middle])
	var part2 = Merge(arr[middle:])
	return merge(part1, part2)
}
func merge[T constraints.Ordered](a []T, b []T) []T {
	var result = make([]T, len(a)+len(b))
	var i = 0
	var j = 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result[i+j] = a[i]
			i++
		} else {
			result[i+j] = b[j]
			j++
		}
	}
	for i < len(a) {
		result[i+j] = a[i]
		i++
	}
	for j < len(b) {
		result[i+j] = b[j]
		j++
	}
	return result
}

// func main() {
// 	arr := []int{2, 3, -4, 4, 1, 8, 9, 0,400,-10,3,50,-1,2}
// 	fmt.Println(MergeSort(arr))
// }
