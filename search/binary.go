package search

// Binary search for target within a sorted array by repeatedly dividing the array in half and comparing the midpoint with the target.
// This function uses recursive call to itself.
// If a target is found, the index of the target is returned. Else the function return -1 and ErrNotFound.
func Binary(array []int, target, low, high int) (int, error) {
	if low > high || len(array) == 0 {
		return -1, ErrNotFound
	}
	mid := int(low + (high-low)/2)
	if array[mid] > target {
		return Binary(array, target, low, mid-1)
	} else if array[mid] < target {
		return Binary(array, target, mid+1, high)
	} else {
		return mid, nil
	}
}

// BinaryIterative search for target within a sorted array by repeatedly dividing the array in half and comparing the midpoint with the target.
// Unlike Binary, this function uses iterative method and not recursive.
// If a target is found, the index of the target is returned. Else the function return -1 and ErrNotFound.
func BinaryIterative(array []int, target int) (int, error) {
	startIndex := 0
	endIndex := len(array) - 1
	var mid int
	for startIndex <= endIndex {

		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] < target {
			startIndex = mid + 1
		} else if array[mid] > target {
			endIndex = mid - 1
		} else {
			return mid, nil
		}
	}
	return -1, ErrNotFound
}
