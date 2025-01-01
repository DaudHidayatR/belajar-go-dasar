package main

import "fmt"

type Person struct {
	Name, address string
	Age           int
}

func main() {
	var p Person
	p.Name = "Daud"
	p.address = "Indonesia"
	p.Age = 23

	fmt.Println(p.Name)
	fmt.Println(p.address)
	fmt.Println(p.Age)
	daud := Person{
		Name:    "Daud",
		address: "Indonesia",
		Age:     23,
	}
	fmt.Println(daud.Name)
	fmt.Println(daud.address)
	fmt.Println(daud.Age)

	bintang := Person{"Bintang", "Indonesia", 23}
	fmt.Println(bintang.Name)
	fmt.Println(bintang.address)
	fmt.Println(bintang.Age)
}
