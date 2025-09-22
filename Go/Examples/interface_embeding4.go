// Embedded Interface in Go

// In Go, an interface can include other interfaces, which is known as interface embedding. This allows one interface to inherit the methods of another, essentially combining multiple interfaces into one. It promotes modular design and makes the code more flexible and scalable.

// Example of Embedded Interface
// Let’s say we are building a system for a smart home.

//We have a Device interface that has methods common to all devices, and we can define more specific interfaces for particular types of devices, like Light and Thermostat.

// Simple Example:

package main

import "fmt"

// Define a Device interface with common methods
type Device interface {
	TurnOn()
	TurnOff()
}

// Define a Light interface that embeds Device and adds its own method
type Light interface {
	Device // Embed the Device interface
	ChangeBrightness(level int)
}

// Define a Thermostat interface that embeds Device and adds its own method
type Thermostat interface {
	Device // Embed the Device interface
	SetTemperature(temp int)
}

// Implement a SmartLight struct
type SmartLight struct {
	name string
}

func (s SmartLight) TurnOn() {
	fmt.Printf("%s is now ON\n", s.name)
}

func (s SmartLight) TurnOff() {
	fmt.Printf("%s is now OFF\n", s.name)
}

func (s SmartLight) ChangeBrightness(level int) {
	fmt.Printf("%s brightness is set to %d\n", s.name, level)
}

// Implement a SmartThermostat struct
type SmartThermostat struct {
	brand string
}

func (t SmartThermostat) TurnOn() {
	fmt.Printf("%s thermostat is now ON\n", t.brand)
}

func (t SmartThermostat) TurnOff() {
	fmt.Printf("%s thermostat is now OFF\n", t.brand)
}

func (t SmartThermostat) SetTemperature(temp int) {
	fmt.Printf("%s thermostat is set to %d degrees\n", t.brand, temp)
}

func main() {
	// Create instances of SmartLight and SmartThermostat
	light := SmartLight{name: "Living Room Light"}
	thermostat := SmartThermostat{brand: "Nest"}

	// Use Light interface for SmartLight
	var l Light = light
	l.TurnOn()
	l.ChangeBrightness(5)
	l.TurnOff()

	fmt.Println()

	// Use Thermostat interface for SmartThermostat
	var t Thermostat = thermostat
	t.TurnOn()
	t.SetTemperature(24)
	t.TurnOff()
}

/*
Advantages of Using Interface Embedding
Reusability of Code:

When multiple interfaces share common functionality (like turning devices on and off), you can embed the shared interface. This reduces redundancy because the methods don’t need to be redefined in every interface.
Modular Design:

Interface embedding makes it easy to build smaller, more focused interfaces, and then combine them into more complex ones. This promotes clean, modular code.
Scalability:

As your system grows, you can easily extend interfaces by embedding more specific behaviors into existing interfaces.
Flexibility:

With interface embedding, any struct that implements the embedded interface can automatically satisfy the parent interface. This gives the system flexibility to evolve without changing existing code structures.
Composability:

Embedding allows you to compose interfaces in a way that reflects the natural structure of your program. For example, a thermostat is a device, so embedding the Device interface in Thermostat makes logical sense and keeps the code intuitive.


Conclusion:
Interface embedding in Go allows you to build powerful, reusable, and modular code. By inheriting methods from other interfaces, you can create more specialized interfaces without duplicating code. This promotes clean and maintainable design patterns in Go.

*/
