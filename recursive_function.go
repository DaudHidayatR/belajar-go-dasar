package main

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result

}
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}
	return value * factorialRecursive(value-1)
}
func main() {
	result := factorialLoop(5)
	re := factorialRecursive(5)
	println(re)
	println(result)
}
