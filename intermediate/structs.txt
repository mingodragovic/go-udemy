package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// method example
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
// method example using pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.Width = r.Width * factor
	r.Height = r.Height * factor
}

func main() {
	// Creating an instance of Person
	p := Person{Name: "Alice", Age: 30}
	fmt.Println("Name:", p.Name, "Age:", p.Age)

	// anonymous struct example
	anon := struct {
		City  string
		State string
	}{
		City:  "New York",
		State: "NY",
	}
	fmt.Println("City:", anon.City, "State:", anon.State)

	r := Rectangle{Width: 5, Height: 10}
	fmt.Println("Area of rectangle:", r.Area())

	// Scaling the rectangle
	r.Scale(3)
	fmt.Println("Scaled rectangle area:", r.Area())

	fmt.Println("-------------")
	nestedStructExample()
	compareStructs()
}

// nested struct example
type Address struct {
	Street string
	City   string
	Zip    string
}
type Employee struct {
	Person  Person
	Address Address
	Role    string
	Company Company

}
type Company struct {
	Name      string
}	
func nestedStructExample() {
	emp := Employee{
		Person: Person{Name: "Bob", Age: 25},
		Address: Address{Street: "123 Main St", City: "Springfield", Zip: "12345"},
		Role: "Developer",
		Company: Company{Name: "TechCorp"},
	}
	fmt.Println("Employee:", emp)
		fmt.Println("Employee street:", emp.Address.Street)
				fmt.Println("Company:", emp.Company.Name)


}
// structure compare example
type Point struct {
	X int
	Y int
}
func compareStructs() {
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 1, Y: 2}
	p3 := Point{X: 2, Y: 3}
	fmt.Println("p1 == p2:", p1 == p2) // true
	fmt.Println("p1 == p3:", p1 == p3) // false
}