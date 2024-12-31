package main

func main() {
	counter := 0
	increment := func() {
		counter++
		println("Increment")
	}
	increment()
	increment()
	println(counter)
}
