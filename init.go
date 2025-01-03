package main

import (
	"belajar-go-dasar/database"
	_ "belajar-go-dasar/internal"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
