package main

import "fmt"

func CetakAsterisk(N int) string {

	var result string
	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i; j++ {
			result += " "
		}
		for k := 1; k <= i; k++ {
			result += "* "
		}
		result += "\n"
	}
	return result

}

func CetakAsteriskBuild(N int) {

	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}

func main() {

	CetakAsteriskBuild(5)
	fmt.Println(CetakAsterisk(5))
}
