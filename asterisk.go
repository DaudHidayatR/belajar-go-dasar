package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	Address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	Address2 := &Address1
	Address2.City = "Bandung"
	fmt.Println(Address1)
	fmt.Println(Address2)

	//Address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	//fmt.Println(Address1)
	//fmt.Println(Address2)
	*Address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(Address1)
	fmt.Println(Address2)
}
