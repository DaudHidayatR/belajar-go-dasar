package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("perulangan ke-", counter)
		counter++
	}
	fmt.Println("perulangan selesai")

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke-", counter)
	}

	names := []string{"andi", "budi", "caca"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("index ke-", index, " = ", name)
	}
}
