package search

import (
	"math"
)

//time O(root(n))
func JumpSearch(arr []int, target, size int) (int, error) {

	if size <= 0 {
		return -1, ErrNotFound
	}
	jumpsize := int(math.Round(math.Sqrt(float64(size))))
	low := 0
	high := jumpsize

	//check high is <= of target or not
	for arr[high] <= target && high < size {
		low = high
		high += int(math.Round(math.Sqrt(float64(size))))
		if high > size {
			high = size
		}
	}
	//linear search in small block
	for i := low; i <= high; i++ {
		if arr[i] == target {
			return i, nil
		}
	}
	return -1, ErrNotFound
}
