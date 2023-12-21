package main

import "fmt"

func main() {
	const firstName string = "Daud"
	const middleName string = "Hidayat"
	const lastName string = "Ramadhan"

	const (
		contact = "08123456789"
		address = "Indonesia"
	)
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
	fmt.Println(contact)
	fmt.Println(address)

	//error
	//firstName = "Bintang"
	//middleName = "Rahmatullah"

}
