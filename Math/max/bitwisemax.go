package max

import "fmt"

func Bitwisemax(a int, b int, base int) int {
	z := a - b
	fmt.Println(z)
	fmt.Println(z >> base)
	fmt.Println((z >> base) & 1)
	i := (z >> base) & 1 // right shift
	fmt.Println((i * z))
	return a - (i * z)
}

// func main() {
// 	fmt.Println(Bitwisemax(32, 64, 63))
// }
