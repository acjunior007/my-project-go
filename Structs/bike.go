package Structs

type Bike struct {
	Make string
	Type string
}

func (b Bike) Drive() string {
	return "Riding a " + b.Type + " bike made by " + b.Make
}
