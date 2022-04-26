package array

import "GoDsAlgo/Math/max"

// kadane's algorithm
func MaximumSubArraySum(input []int) int {
	if len(input) <= 0 {
		return 0
	}
	var currentMax, globalMax = input[0], input[0]

	for i := 1; i < len(input)-1; i++ {
		if currentMax < 0 {
			currentMax = input[i]
		} else {
			currentMax += input[i]
		}
		if globalMax < currentMax {
			globalMax = currentMax
		}
	}
	return globalMax
}

func MaxSubarraySum(array []int) int {
	var currentMax int
	var maxTillNow int
	if len(array) != 0 {
		currentMax = array[0]
		maxTillNow = array[0]
	}
	for _, v := range array {
		currentMax = max.Int(v, currentMax+v)
		maxTillNow = max.Int(maxTillNow, currentMax)
	}
	return maxTillNow
}
