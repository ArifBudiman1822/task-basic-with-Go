package main

import "fmt"

func main() {

	var bilangan int
	fmt.Print("Masukkan Bilangan: ")
	fmt.Scanf("%d", &bilangan)

	for pembagi := bilangan; pembagi <= bilangan; pembagi-- {
		if pembagi == 0 {
			break
		}
		if bilangan%pembagi == 0 {
			fmt.Println(pembagi)
		}
	}
}
