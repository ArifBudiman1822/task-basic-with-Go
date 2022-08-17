package main

import "fmt"

func UbahHuruf(sentence string) string {

	// ASCII A-Z = urutan huruf == 65 - 90
	var enkripsi string //input enkripsi
	// ambil index sentence[i] jadikan int//ASCII
	for i := 0; i < len(sentence); i++ {
		urutanhurufFromAscii := int(sentence[i])
		// digeser 10 huruf ke kanan
		urutanhurufFromAscii = urutanhurufFromAscii + 10
		// diubah ASCII ke string(rune(sentence[i]))
		if string(rune(sentence[i])) == " " {
			enkripsi += string(rune(sentence[i]))
		} else if urutanhurufFromAscii > 90 {
			urutanhurufFromAscii = urutanhurufFromAscii - 26
			enkripsi += string(rune(urutanhurufFromAscii))
		} else {
			enkripsi += string(rune(urutanhurufFromAscii))
		}
	}
	return enkripsi

}

func main() {

	fmt.Println(UbahHuruf(" "))
	fmt.Println(UbahHuruf("SEPULSA OKE"))
	fmt.Println(UbahHuruf("ALTERRA ACADEMY"))
	fmt.Println(UbahHuruf("INDONESIA"))
	fmt.Println(UbahHuruf("GOLANG"))
	fmt.Println(UbahHuruf("PROGRAMMER"))
}
