package main

import "fmt"

func main() {
	//person := map[string]string{}
	//person["name"] = "daud"
	//person["address"] = "yogyakarta"
	//person["job"] = "programmer backend"

	person := map[string]string{
		"name":    "daud",
		"address": "yogyakarta",
		"job":     "programmer backend",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["job"])

	book := make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "daud"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "ups")
	fmt.Println(book)

}
