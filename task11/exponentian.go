package main

import "fmt"

func pangkatBuild(base, pangkat int) (hasil int) {

	if pangkat == 0 {
		return 1
	} else if pangkat == 1 {
		return base
	} else if pangkat == 2 {
		return base * base
	}

	hasil = base * base

	for i := 2; i < pangkat; i++ {
		hasil *= base
	}
	return hasil
}

func main() {

	fmt.Println(pangkatBuild(2, 3))
	fmt.Println(pangkatBuild(5, 3))
}
