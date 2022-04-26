package sort

import (
	"GoDsAlgo/Math/max"
)

func countingSort(arr []int, exp int) []int {
	//fmt.Println("%d exp", exp)
	var digits [10]int // for number 10, for char 26
	var output = make([]int, len(arr))

	for _, item := range arr {
		digits[(item/exp)%10]++
	}
	//fmt.Println(digits)
	for i := 1; i < 10; i++ {
		digits[i] += digits[i-1]
	}

	//fmt.Println(digits)
	for i := len(arr) - 1; i >= 0; i-- {
		output[digits[(arr[i]/exp)%10]-1] = arr[i]
		digits[(arr[i]/exp)%10]--
	}

	// fmt.Println(output)
	// fmt.Println("----------------")
	return output
}

func unsignedRadixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	maxElement := max.Int(arr...)
	for exp := 1; maxElement/exp > 0; exp *= 10 {
		arr = countingSort(arr, exp)
	}
	return arr
}

func RadixSort(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}

	var negatives, nonNegatives []int

	for _, value := range arr {
		if value < 0 {
			negatives = append(negatives, -value)
		} else {
			nonNegatives = append(nonNegatives, value)
		}
	}
	negatives = unsignedRadixSort(negatives)
	//fmt.Println(negatives)
	// // Reverse the negative array and restore signs
	for i, j := 0, len(negatives)-1; i <= j; i, j = i+1, j-1 {
		negatives[i], negatives[j] = -negatives[j], -negatives[i]
	}
	return append(negatives, unsignedRadixSort(nonNegatives)...)
}

// func main() {
// 	//arr := []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10}
// 	arr := []int{1, 8, 9, 0, 0, 1, 1, 1, 2, 40}
// 	fmt.Println(RadixSort(arr))
// }
