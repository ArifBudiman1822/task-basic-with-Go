package main

import "fmt"

func bilKeren(bil int) {
	jatahpembagi := 0
	for pembagi := 1; pembagi <= bil; pembagi++ {
		if (bil%1 == 0) && (bil == pembagi) {
			fmt.Println(bil, "adalah bilangan keren")
			break
		}
		if (bil == pembagi) || (jatahpembagi > 4) {
			fmt.Println(bil, "bukan bilangan keren")
			break
		}
		if bil%pembagi == 0 {
			jatahpembagi++
		}
	}

}

func main() {
	// input bil
	//bil keren = bil % 1 == 0 && bil == pembagi dan habis dibagi 2 bil postif
	// bil keren jatahpembagi <= 4
	// bukan bilkeren bil = pembagi || jatahpembagi > 4
	// jika bil % pembagi == 0 maka jatah++

	bilKeren(17)
	bilKeren(51)
	bilKeren(52)

}
