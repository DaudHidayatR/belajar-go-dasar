package main

import "fmt"

func getCompleateName() (firstName string, middleName string, lastName string) {
	firstName = "John"
	middleName = "Doe"
	lastName = "Smith"
	return firstName, middleName, lastName

}
func main() {
	firstName, middleName, lastName := getCompleateName()
	fmt.Println(firstName, middleName, lastName)

}
