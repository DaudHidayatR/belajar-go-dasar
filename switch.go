package main

import "fmt"

func main() {
	var i int = 10
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	case 4:
		fmt.Println("i is 4")
	case 5:
		fmt.Println(("i is 5"))
	default:
		fmt.Println(("i is not 1-5"))
	}
	name := "daud"
	switch langth := len(name); langth > 5 {
	case true:
		fmt.Println("name is true")
	case false:
		fmt.Println("name is false")
	}

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("name is true")
	case length < 5:
		fmt.Println("name is false")
	default:
		fmt.Println("name is 5")
	}

}
