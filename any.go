package main

import "fmt"

func Ups() any {
	return 1
}
func main() {
	var data any = Ups()
	fmt.Println(data)
}
