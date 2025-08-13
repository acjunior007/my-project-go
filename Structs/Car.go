package Structs

import (
	"fmt"
)

type Car struct {
	Brand string
	Model string
	Year  int
	Engine
	Transmission
}

func InitializeCar(brand, model string, year int) Car {
	return Car{
		Brand: brand,
		Model: model,
		Year:  year,
	}
}

func InitializeCarEngine(brand, model string, year int, engine Engine) Car {
	return Car{
		Brand:  brand,
		Model:  model,
		Year:   year,
		Engine: engine,
	}
}

func InitializeCarComplete(brand, model string, year int, engine Engine, transmission Transmission) Car {
	return Car{
		Brand:        brand,
		Model:        model,
		Year:         year,
		Engine:       engine,
		Transmission: transmission,
	}
}

func (c Car) GetCarInfo() string {
	return fmt.Sprintf("Car Info: %s %s, Year: %d", c.Brand, c.Model, c.Year)
}

// Method with pointer receiver
//* Using a pointer receiver allows you to modify the original struct instance directly
//* When you need to modify the struct`s fields, prefer pointer receivers

func (c *Car) UpdateYearPointer(year int) {
	c.Year = year
}

func (c Car) Drive() string {
	return fmt.Sprintf("Driving a %d %s %s", c.Year, c.Brand, c.Model)
}
