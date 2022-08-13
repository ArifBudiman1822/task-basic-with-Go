package main

import (
	"fmt"
	"strconv"
)

func bilPrima(N int) bool {
	// jika N <= 1 maka tidak prima
	if N <= 1 {
		return false
	}

	//bil prima == hanya bisa dibagi 1 dan N(bilangannya sendiri)
	// tentukan kebalikannya
	for pembagi := 2; pembagi < N; pembagi++ {
		if N%pembagi == 0 {
			return false
		}
	}
	// jika N ! yang diatas maka return true
	return true

}

func fullPrima(N int) bool {

	//jika func bilprima adalah true maka !(kebalikannya) false
	if !bilPrima(N) {
		return false
	}

	// ubah int ke string "strconv.Itoa(N)"
	strN := strconv.Itoa(N)

	// buat index untuk mengambil 1 persatu karakter dari string "strN"
	for i := 0; i < len(strN); i++ {
		// ubah string "strN" ke int "intN"
		intN, _ := strconv.Atoi(string(strN[i]))
		// ambil intN ke funcbilprima
		// jika !bilprima(intN) == return false
		if !bilPrima(intN) {
			return false
		}
	}
	//! yang diatas
	return true

}

func main() {

	fmt.Println(bilPrima(17))
	fmt.Println(fullPrima(2))
	fmt.Println(fullPrima(5))
	fmt.Println(fullPrima(13))
	fmt.Println(fullPrima(17))
}
