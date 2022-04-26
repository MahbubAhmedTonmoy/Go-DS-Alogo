package sort

import (
	"GoDsAlgo/constraints"
)

func ShellSort[T constraints.Ordered](arr []T) []T {
	n := len(arr)
	for gap := n / 2; gap >= 1; gap /= 2 {
		for j := gap; j < n; j++ {
			for i := j - gap; i >= 0; i = i - gap {
				if arr[i+gap] > arr[i] {
					break
				} else {
					arr[i+gap], arr[i] = arr[i], arr[i+gap]
					//this check previous elements after 1 sawp
				}
			}
		}
	}
	return arr
}

// func main() {
// 	arr := []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10}
// 	//arr = []int{1,8,9,0 ,0,1,1,1,2,40,-1}
// 	fmt.Println(Shell(arr))
// }
