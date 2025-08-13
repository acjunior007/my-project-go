package Structs

import "fmt"

type Engine struct {
	Horsepower int
}

func (e Engine) Start() {
	fmt.Println("Engine started with", e.Horsepower, "horsepower")
}
