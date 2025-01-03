package helper

import "fmt"

var version = "1.0.0"
var Application = "Golang"

func sayGoodBye() {
	fmt.Println("Good Bye")
}

func SayHello(name string) string {
	return "Hello " + name
}
