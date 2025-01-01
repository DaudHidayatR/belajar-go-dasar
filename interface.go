package main

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	println("Hello", value.GetName())
}

type person struct {
	Name string
}

func (person person) GetName() string {
	return person.Name

}

type animal struct {
	Name string
}

func (animal animal) GetName() string {
	return animal.Name

}

func main() {
	var daud person
	daud.Name = "Daud"
	SayHello(daud)

	var kucing animal
	kucing.Name = "Kucing"
	SayHello(kucing)

}
