package main

import "fmt"

func main() {
	// input bil baris ke N
	// buat perulangan pembagi hingga ke N
	// jika N % pembagi == 0 maka lampu menyala dan sebaliknya
	// input checktombol hidup == ganjil, mati == genap

	var N int
	N = 5
	check := 0

	for pembagi := 1; pembagi <= N; pembagi++ {
		if N%pembagi == 0 {
			check++
		}
	}

	if check%2 == 0 {
		fmt.Println(N, "Lampu Mati")
	} else {
		fmt.Println(N, "Lampu Menyala")
	}

}
