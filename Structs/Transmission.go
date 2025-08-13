package Structs

import "fmt"

type Transmission struct {
	Type string
}

func (t Transmission) Shift() {
	fmt.Println("Shifting to:", t.Type, "transmission.")
}
