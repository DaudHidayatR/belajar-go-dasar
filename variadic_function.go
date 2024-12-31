package main

import "fmt"

func sumAll(Numbers ...int) int {
	sum := 0
	for _, number := range Numbers {
		sum += number
	}
	return sum
}
func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5))
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11))
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12))

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sumAll(numbers...))
}
