package main

import "fmt"

func main() {
	name := "daud"
	if name == "daud" {
		fmt.Println("Hello Daud")
	} else if name == "sandi" {
		fmt.Println("Hello Sandi")
	} else {
		fmt.Println("Hello Stranger")
	}
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
