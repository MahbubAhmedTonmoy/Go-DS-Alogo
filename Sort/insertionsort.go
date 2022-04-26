package sort

import (
	"GoDsAlgo/constraints"
)

func Insertionsort[T constraints.Ordered](arr []T) []T {

	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i
		for ; j > 0 && arr[j-1] >= temp; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
		//arr[j] = temp
	}
	return arr
}

// func main() {
// 	arr := []int{2, 3, -4, 4, 1, 8, 9, 0}
// 	fmt.Println(Insertionsort(arr))
// }
