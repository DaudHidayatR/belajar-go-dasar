package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	var d = 10
	var e = 10
	var f = d - e
	fmt.Println(f)

	var g = 10
	var h = 10
	var i = g * h
	fmt.Println(i)

	var j = 10
	j += 10
	fmt.Println(j)

	j++
	fmt.Println(j)

	j--
	fmt.Println(j)

}
