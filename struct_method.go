package main

import "fmt"

type customType struct {
	name, address string
	age           int
}

func (customType customType) printName(name string) {
	fmt.Println("Name:", name, "my name is", customType.name)
}

func main() {
	var p customType
	p.name = "Daud"
	p.address = "Indonesia"
	p.age = 23

	fmt.Println(p.name)
	fmt.Println(p.address)
	fmt.Println(p.age)

	daud := customType{
		name:    "Daud",
		address: "Indonesia",
		age:     23,
	}
	fmt.Println(daud.name)
	fmt.Println(daud.address)
	fmt.Println(daud.age)

	bintang := customType{"Bintang", "Indonesia", 23}
	fmt.Println(bintang.name)
	fmt.Println(bintang.address)
	fmt.Println(bintang.age)

	p.printName("Daud")
	daud.printName("Daud")
	bintang.printName("Daud")
}
