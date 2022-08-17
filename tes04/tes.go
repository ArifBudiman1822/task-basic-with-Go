package main

import "fmt"

func TabelPerkalian(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println("")
	}
}

func main() {

	TabelPerkalian(9)
}
