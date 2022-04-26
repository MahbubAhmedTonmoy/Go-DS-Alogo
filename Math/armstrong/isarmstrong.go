package armstrong

import (
	"math"
	"strconv"
)

// 153
// 1^3 + 5^3 + 3^3 = 153
func isarmstrong(number int) bool {
	var rigthMost int
	var sum = 0
	var tempNum int = number

	length := float64(len(strconv.Itoa(number)))
	for tempNum > 0 {
		rigthMost = tempNum % 10
		sum += int(math.Pow(float64(rigthMost), length))
		tempNum /= 10
	}
	return number == sum
}
