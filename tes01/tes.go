package main

import (
	"fmt"
	"strconv"
)

func RumusBilPrime(N int) bool {

	if N <= 1 {
		return false
	}

	for i := 2; i < N; i++ {
		if N%i == 0 {
			return false
		}
	}

	return true
}

func FullPrime(N int) bool {
	if !RumusBilPrime(N) {
		return false
	}
	// setiap digit N di uji func bil prima
	strAngka := strconv.Itoa(N)
	for i := 0; i < len(strAngka); i++ {
		intString, _ := strconv.Atoi(string(strAngka[i]))
		if !RumusBilPrime(intString) {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(FullPrime(2))
	fmt.Println(FullPrime(3))
	fmt.Println(FullPrime(17))
	fmt.Println(FullPrime(23))
	fmt.Println(FullPrime(11))
	fmt.Println(FullPrime(13))
}
