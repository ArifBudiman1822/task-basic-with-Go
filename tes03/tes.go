package main

import "fmt"

func CetakXYZ(N int) {

	// Y == ganjil
	// Z == genap
	// X == % 3 == 0
	var check int
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			check++
			if check%3 == 0 {
				fmt.Print("X ")
			} else if check%2 == 0 {
				fmt.Print("Z ")
			} else {
				fmt.Print("Y ")
			}
		}
		fmt.Println("")
	}

}

func CetakXYZBuild(N int) string {

	var result string
	var check int

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			check++
			if check%3 == 0 {
				result += "X "
			} else if check%2 == 0 {
				result += "Z "
			} else {
				result += "Y "
			}
		}
		result += "\n"
	}
	return result

}

func main() {

	CetakXYZ(5)
	fmt.Println("\n")
	fmt.Println(CetakXYZBuild(5))
}
