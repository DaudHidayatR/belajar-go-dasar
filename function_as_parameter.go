package main

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	println("Hello", nameFiltered)
}
func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}

}

func main() {

	sayHelloWithFilter("daud", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
}
