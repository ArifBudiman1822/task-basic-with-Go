package main

import "fmt"

func gradeScore(studentScore int) {

	if studentScore >= 100 || studentScore <= 0 {
		fmt.Println("Unknown")
	} else if studentScore >= 80 {
		fmt.Println(studentScore, "= Nilai A")
	} else if studentScore >= 65 {
		fmt.Println(studentScore, "= Nilai B+")
	} else if studentScore >= 50 {
		fmt.Println(studentScore, "= Nilai B")
	} else if studentScore >= 35 {
		fmt.Println(studentScore, "= Nilai C")
	} else if studentScore >= 0 {
		fmt.Println(studentScore, "= Nilai D")
	}
}

func main() {

	gradeScore(80)
	gradeScore(110)
	gradeScore(-20)
}
