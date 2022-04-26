package search

import "fmt"

// func InterpolationSort(sortedArr []int, terget int) (int, error) {

// 	if len(sortedArr)-1 == 0 {
// 		return -1, ErrNotFound
// 	}

// 	var (
// 		low, high           = 0, len(sortedArr) - 1
// 		lowValue, highValue = sortedArr[low], sortedArr[high]
// 	)
// 	for lowValue != highValue && (lowValue <= terget) && (highValue >= terget) {
// 		mid := low + int(float64(float64((terget-lowValue)*(high-low))/float64(highValue-lowValue)))
// 		if sortedArr[mid] == terget {
// 			for mid > 0 && sortedArr[mid-1] == terget {
// 				mid--
// 			}
// 			return mid, nil
// 		}
// 		if sortedArr[mid] > terget {
// 			high, highValue = mid-1, sortedArr[high]
// 		} else {
// 			low, lowValue = mid+1, sortedArr[low]
// 		}
// 	}
// 	if terget == low {
// 		return low, nil
// 	}
// 	return -1, ErrNotFound
// }

func InterpolationSort(array []int, target, low, high int) (int, error) {

	fmt.Println("arr ,value, low, high", array, target, low, high)
	if len(array) == 0 {
		return -1, ErrNotFound
	}
	mid := low + int(float64(float64((target-array[low])*(high-low))/float64(array[high]-array[low])))

	if low > high || mid < low || mid > high {
		return -1, ErrNotFound
	}

	fmt.Println("mid", mid)
	if array[mid] > target {
		return InterpolationSort(array, target, low, mid-1)
	} else if array[mid] < target {
		return InterpolationSort(array, target, mid+1, high)
	} else {
		return mid, nil
	}
}
