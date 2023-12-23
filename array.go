package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Daud"
	name[1] = "Hidayat"
	name[2] = "Ramadhan"

	fmt.Println(name)
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [...]int{
		90,
		95,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])

	values[0] = 100
	fmt.Println(len(values))

}
