package main

import "fmt"

func main() {

	// buat gambar asterisk dari for
	//    	*       bintang menggandakan diri sampai nilai ke N
	//     * *
	//    * * *
	// 	 * * * *
	//	* * * * *

	var N int = 5

	// pengulangan index ke N
	for i := 1; i <= N; i++ {
		// pengulangan spasi dengan J <= N - i
		for j := 1; j <= N-i; j++ {
			fmt.Print(" ")
		}
		// pengulangan cetak "* " k <= i
		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}
		//buat baris baru dengan println
		fmt.Println("")
	}
}
