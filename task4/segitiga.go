package main

import "fmt"

func luasSegitiga(alas, tinggi float64) {
	fmt.Println(alas * tinggi / 2)
}

func main() {
	luasSegitiga(20, 25)
}
