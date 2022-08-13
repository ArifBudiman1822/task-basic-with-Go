package main

import "fmt"

func palindrome(input string) bool {
	// hitung panjang input
	// bandingkan huruf input awal dengan input huruf terakhir
	// jika sama maka return true

	for i := 0; i < len(input)/2; i++ {
		if input[i] == input[len(input)-1-i] {
			return true
		}
	}
	return false
}

func main() {

	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("kupu kupu"))

}
