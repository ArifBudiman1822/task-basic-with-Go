package main

import (
	"fmt"
)

func DrawXYZ(N int) string {

	// kelipatan 3 == Z
	// genap == X
	// ganjil == Y
	var check int
	var result string
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			check++
			if check%3 == 0 {
				result += "Z "
			} else if check%2 == 0 {
				result += "X "
			} else {
				result += "Y "
			}
		}
		result += "\n"
	}
	return result
}

func main() {
	fmt.Print(DrawXYZ(3), "\n")
	fmt.Print(DrawXYZ(5), "\n")
	fmt.Print(DrawXYZ(1), "\n")

}
