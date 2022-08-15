package main

import (
	"fmt"
	"strconv"
)

func bilPrima(N int) bool {
	if N <= 1 {
		return false
	}
	// bil yang habis dibagi 1 dan bilangannya sendiri adalah bilprima
	for i := 2; i < N; i++ {
		if N%i == 0 {
			return false
		}
	}
	return true

}

func fullPrima(N int) bool {
	// uji N dengan func bilPrima kebalikannya
	if !bilPrima(N) {
		return false
	}

	// ubah N ke string
	stringAngka := strconv.Itoa(N)
	// ambil index string angka uji dengan func satu satu
	for i := 1; i < len(stringAngka); i++ {
		angka, _ := strconv.Atoi(string(stringAngka[i]))
		if !bilPrima(angka) {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(fullPrima(3))
	fmt.Println(fullPrima(23))
	fmt.Println(fullPrima(22))
}
