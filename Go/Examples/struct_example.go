package main

import "fmt"

type Player struct {
	name  string
	age   int
	state string
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Counter struct {
	Count int
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c *Counter) Increment() {
	c.Count++
}

type Address struct {
	City  string
	State string
}

type Employee struct {
	Name    string
	Age     int
	Address // Nested struct
}

type User struct {
	Person
	Role string
}

type Admin struct {
	Person
	AccessLevel int
}

type Person struct {
	Name  string
	Email string
}

func main() {

	//Using feild name
	ply1 := Player{
		name:  "Omkar",
		age:   24,
		state: "Karnataka",
	}

	var ply2 Player
	ply2 = Player{"Rutuja", 23, "Maharashtra"}

	fmt.Printf("Data type of Player 1 is %T and value is: %+v\n", ply1, ply1)
	fmt.Printf("Data type of Player 2 is %T and value is: %+v\n", ply2, ply2)
	fmt.Printf("To print player name only : %v\n", ply1.name)

	area := Rectangle{Length: 5.0, Width: 3.0}
	fmt.Println("Area of rectangle :", area.Area())

	counter := Counter{Count: 0}
	counter.Increment()
	fmt.Println("Counter after increment :", counter.Count)
	counter.Increment()
	fmt.Println("Counter after 2nd increment :", counter.Count)

	emp := Employee{
		Name: "Omi",
		Age:  24,
		Address: Address{
			City:  "Sankeshwar",
			State: "Karnataka",
		},
	}

	fmt.Println("Emp name :", emp.Name)
	fmt.Println("Emp City :", emp.Address.City)

	user := User{
		Person: Person{"SPS", "sps@cuda.com"},
		Role:   "Soft Engg",
	}

	admin := Admin{
		Person: Person{
			Name:  "Admin User",
			Email: "admin@cuda.com",
		},
		AccessLevel: 1,
	}

	fmt.Printf("User: %+v\n", user)
	fmt.Printf("Admin: %+v\n", admin)
}
