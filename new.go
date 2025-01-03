package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat1 = new(Address)
	var alamat2 = alamat1

	alamat2.City = "Bandung"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
