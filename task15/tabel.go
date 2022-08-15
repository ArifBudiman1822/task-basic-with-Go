package main

import "fmt"

func tabelPerkalian(number int) {
	// buat perulangan big O notation 9 * 9
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			hasil := i * j
			fmt.Print(hasil, " ")
		}
		fmt.Println("")
	}
}

func main() {
	tabelPerkalian(9)
}
