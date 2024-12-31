package main

import "fmt"

func getFullname() (string, string) {
	return "John", "Doe"

}
func main() {
	firstname, _ := getFullname()
	fmt.Println(firstname)
}
