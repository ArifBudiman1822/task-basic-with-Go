package main

import "fmt"

func bilPrima(bil int) bool {
	//input bl
	//jika bil <= 1 maka false
	// jika genap maka false jika ganjil maka true
	// pengulangan pembagi dari 2 sampai ke bill
	// jika ada yang habis dibagi maka false

	if bil <= 1 {
		return false
	}

	for pembagi := 2; pembagi < bil; pembagi++ {
		if bil%pembagi == 0 {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(bilPrima(17))
	fmt.Println(bilPrima(13))
	fmt.Println(bilPrima(52))
	fmt.Println(bilPrima(35))
}
