package main

import "fmt"

const phi = 3.14

func luasTabung(r, tinggi float64) float64 {
	return 2 * phi * r * (r + tinggi)
}

func main() {
	fmt.Println(luasTabung(4, 20))

	//var r float64 = 4
	//var tinggi float64 = 20

	//luasTabung := 2 * phi * r * (r + tinggi)
	//fmt.Println(luasTabung)

}
