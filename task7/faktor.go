package main

import "fmt"

func main() {
	//angka := 89
	//text := "lalalala"
	//angka13 := 13
	//tes := fmt.Sprintf("%v-%v\t%v", angka, text, angka13)
	//fmt.Println(tes)
	var bilangan int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanf("%d", &bilangan)

	for pembagi := 1; pembagi <= bilangan; pembagi++ {
		if bilangan%pembagi == 0 {
			fmt.Println(pembagi)
		}
	}
}
