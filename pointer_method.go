package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}
func main() {
	daud := Man{"Daud"}
	daud.Married()
	fmt.Println(daud.Name)
}
