package Structs

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}
