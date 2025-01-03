package main

import "fmt"

func random() any {
	return "ok"
}
func main() {
	var result = random()
	//var resultString string = result.(string)
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	}
	fmt.Println(result)

}
