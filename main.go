package main

import (
	"fmt"
	interfaces "my-project-go/Interfaces"
	"my-project-go/Structs"
)

func main() {

	// fmt.Println("Hello, World!")
	//
	//	for i := range 4 {
	//		parOuImpar(i)
	//	}
	// res := add(3, 2)
	// fmt.Println("Resultado: ", res)

	//* Par ou Ímpar (short condition)
	// if num := -1; num%2 == 0 {
	// 	fmt.Println("Even : ", num)
	// } else {
	// 	fmt.Println("Odd : ", num)
	// }

	//* Verificação de idade
	// age := 17
	// hasID := true

	// if age >= 18 && hasID {
	// 	fmt.Println("You can enter the club.")
	// } else {
	// 	fmt.Println("You cannot enter the club.")
	// }

	//* Switch
	// day := "Friday"
	// switch day {
	// case "Monday":
	// 	fmt.Println("Start the workweek!")
	// case "Friday":
	// 	fmt.Println("Weekend is near!")
	// case "Saturday", "Sunday":
	// 	fmt.Println("Enjoy the weekend!")
	// default:
	// 	fmt.Println("A regular day!")
	// }

	//* Switch with fallthrough
	// rating := 4
	// switch rating {
	// case 5:
	// 	fmt.Println("Excellent!")
	// 	fallthrough // This will execute the next case as well
	// case 4:
	// 	fmt.Println("Very Good!")
	// case 3:
	// 	fmt.Println("Good!")
	// 	fallthrough // This will execute the next case as well
	// default:
	// 	fmt.Println("Poor!")
	// }

	//* Switch with interface
	// checkType(42)
	// checkType("Hello")
	// checkType(true)
	// checkType(3.14)

	//* For statement
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Iteration:", i)
	// }

	//* For as while - go just have the statment for
	// count := 0
	// for count < 5 {
	// 	fmt.Println("Count:", count)
	// 	count++
	// }

	//* Looping over collection
	// number := []int{10, 20, 30, 40, 50}

	// for index, value := range number {
	// 	fmt.Println("Index:", index, "Value:", value)
	// }

	//* Lopping over a map
	// person := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
	// for name, age := range person {
	// 	fmt.Println("Name:", name, "Age:", age)
	// }

	//* Returning multiples values
	//q, r := divide(16, 3)
	//fmt.Println("Quotient:", q, "Remainder:", r)

	//* Ignoring the return value
	//q, _ := divide(16, 3) //! ignoring remainder
	//fmt.Println("Quotient:", q)

	//* Named Return Values
	// s, p := calculate(3, 4)
	// fmt.Println("Sum:", s, "Product:", p)

	//* Functions to many parameters
	//total := sum(25, 2)
	//fmt.Println("Total sum:", total)

	//* Slices
	// numbs := []int{4, 4}
	// total := sum(numbs...)
	// fmt.Println("Total sum:", total)

	//* Regular and variadic parameters
	//greetUsers("Hello", "Alice", "Bob", "Charlie")
	//* example 01
	// add := func(a, b int) int {
	// 	return a + b
	// }
	// res := add(10, 4)
	// fmt.Println("Resultado:", res)

	// //* example 02
	// result := func(a, b int) int {
	// 	return a * b
	// }(5, 4)
	// fmt.Println("Resultado:", result)

	// nextCount := counter()
	// fmt.Println("Next Count:", nextCount()) //? 1
	// fmt.Println("Next Count:", nextCount()) //? 2
	// fmt.Println("Next Count:", nextCount()) //? 3

	//* Return errors
	// result, error := divideWithError(15, 3)
	// if error != nil {
	// 	fmt.Println("Error:", error)
	// } else {
	// 	fmt.Println("Result:", result)
	// }

	//* Handling errors in file operations
	// file, err := os.Open("noexistfile.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("File opened successfully:", file.Name())

	//* Custom error handling
	// err := validateAge(18)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Age is valid.")
	// }

	//* Error handling with multiple errors
	// err := Process()
	// if err != nil {
	// 	// The error returned by errors.Join satisfies this interface.
	// 	// We can check for it to handle the errors individually.
	// 	type unwrapper interface {
	// 		Unwrap() []error
	// 	}

	// 	if e, ok := err.(unwrapper); ok {
	// 		fmt.Println("Multiple errors occurred:")
	// 		for _, wrappedErr := range e.Unwrap() {
	// 			fmt.Println(" -", wrappedErr)
	// 		}
	// 	} else {
	// 		// Fallback for a single error that doesn't unwrap into a slice.
	// 		fmt.Println("An error occurred:", err)
	// 	}
	// }

	//* Checking for a specific error
	// err := getResource(0)
	// if errors.Is(err, ErrNotFound) {
	// 	fmt.Println("Resource not found error occurred!")
	// } else {
	// 	fmt.Println("Other error occurred:", err)
	// }

	//* Initializing a struct
	// car := Struct.InitializeCar("Hyundai", "Creta", 2020)
	// fmt.Println("Before:", car.GetCarInfo())
	// fmt.Println("----------------------------------")

	// car.UpdateYearPointer(2025)
	// fmt.Println("After update by pointer:", car.GetCarInfo())

	//* Cart with embedded struct
	//myCar := Structs.InitializeCarEngine("Hyundai", "Creta", 2020, Structs.Engine{Horsepower: 150})
	// myCar := Structs.InitializeCarComplete("Hyundai", "Creta", 2020, Structs.Engine{Horsepower: 150}, Structs.Transmission{Type: "Manual"})
	// fmt.Println("Car Info:", myCar.GetCarInfo())
	// myCar.Start()
	// myCar.Shift()

	//* Working with interfaces
	// Create a person instance
	// person := Structs.Person{Name: "Alice"}
	// // Use the Speak method defined in the Speaker interface
	// var s interfaces.Speaker = person
	// fmt.Println(s.Speak())

	//* Interface with polimorphism
	// Create a instance of circle and retangle
	// circle := Structs.Circle{Radius: 5}
	// rectangle := Structs.Rectangle{Length: 10, Width: 5}

	// Pass the both types to the printArea method
	// printArea(circle)
	// printArea(rectangle)

	//* Polymorphism with interfaces
	car := Structs.Car{Brand: "Toyota", Model: "Camry", Year: 2020}
	bike := Structs.Bike{Make: "Yamaha", Type: "Sport"}

	// Use polimorphism to handle different types
	StartVehicle(car)
	StartVehicle(bike)

} //* end main

// * Function that accepts the Vehicle interface
func StartVehicle(v interfaces.Vehicle) {
	fmt.Println(v.Drive())
}

//* Polymorphism example
// func printArea(s polimorphism.Shape) {
// 	fmt.Println("Area:", s.Area())
// }

// func parOuImpar(i int) {
// 	if i%2 == 0 {
// 		fmt.Println("par: ", i)
// 	} else {
// 		fmt.Println("impar: ", i)
// 	}
// }

// // * Função de adição
// func add(a int, b int) int {
// 	return a + b
// }

// // * Verificação de tipo
// func checkType(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Println("Integer:", v)
// 	case string:
// 		fmt.Println("String:", v)
// 	case bool:
// 		fmt.Println("Boolean:", v)
// 	default:
// 		fmt.Println("Unknown type")
// 	}
// }

// // * Returning multiple values
// func divide(a, b int) (int, int) {
// 	quotient := a / b
// 	remainder := a % b
// 	return quotient, remainder
// }

// func calculate(a, b int) (sum int, product int) {
// 	sum = a + b
// 	product = a * b
// 	return
// }

// func sum(numbers ...int) int {
// 	total := 0
// 	for _, n := range numbers {
// 		total += n
// 	}
// 	return total
// }

// func greetUsers(prefix string, names ...string) {
// 	for _, name := range names {
// 		fmt.Printf("%s, %s!\n", prefix, name)
// 	}
// }

// func counter() func() int {
// 	count := 0
// 	return func() int {
// 		count++ //! increment the counter
// 		return count
// 	}
// }

// func divideWithError(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// func validateAge(age int) error {
// 	if age < 18 {
// 		return errors.New("Age must be at least 18")
// 	}
// 	return nil
// }
