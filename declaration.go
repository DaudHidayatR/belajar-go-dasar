package main

import "fmt"

func main() {
	type NoKTP string

	var noKtpDaud NoKTP = "1234567890"

	var noKtpBintang string = "1234567890"
	var contohKTP NoKTP = NoKTP(noKtpBintang)

	fmt.Println(noKtpDaud)
	fmt.Println(contohKTP)

}
