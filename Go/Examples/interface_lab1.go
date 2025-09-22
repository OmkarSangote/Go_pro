/*
Lab 1: Smart Home Devices

Problem Statement:
You are working on a smart home system where different types of devices need to be controlled. All devices should have the ability to turn on and off, and some devices may have additional functionalities. You need to create a Device interface for basic controls, and then embed this interface into more specialized interfaces for specific types of devices.

Define a Device interface that has methods TurnOn() and TurnOff().

Create a Light interface that embeds the Device interface and adds a method ChangeBrightness(level int).

Create a Speaker interface that embeds the Device interface and adds a method PlayMusic(song string).

Implement the SmartLight and SmartSpeaker structs with appropriate methods.

In the main() function, create instances of SmartLight and SmartSpeaker, and control them using their respective interfaces.
*/

//Skeleton Code:

//package main

// Define the Device interface

// Define the Light interface that embeds Device and adds ChangeBrightness

// Define the Speaker interface that embeds Device and adds PlayMusic

// Define the SmartLight struct

// Define the SmartSpeaker struct

// Implement methods for SmartLight

// Implement methods for SmartSpeaker

//func main() {
// Create instances of SmartLight and SmartSpeaker

// Use the Light interface for controlling SmartLight

// Use the Speaker interface for controlling SmartSpeaker
// }
package main

import "fmt"

// Define the Device interface
type Device interface {
	TurnOn()
	TurnOff()
}

// Define the Light interface that embeds Device and adds ChangeBrightness
type Light interface {
	Device
	ChangeBrightness(level int)
}

// Define the Speaker interface that embeds Device and adds PlayMusic
type Speaker interface {
	Device
	PlayMusic(song string)
}

// Define the SmartLight struct
type SmartLight struct {
	name string
}

// Define the SmartSpeaker struct
type SmartSpeaker struct {
	brand string
}

// Implement methods for SmartLight
func (s SmartLight) TurnOn() {
	fmt.Printf("%s light is now ON\n", s.name)
}

func (s SmartLight) TurnOff() {
	fmt.Printf("%s light is now OFF\n", s.name)
}

func (s SmartLight) ChangeBrightness(level int) {
	fmt.Printf("%s light brightness is set to %d\n", s.name, level)
}

// Implement methods for SmartSpeaker
func (s SmartSpeaker) TurnOn() {
	fmt.Printf("%s speaker is now ON\n", s.brand)
}

func (s SmartSpeaker) TurnOff() {
	fmt.Printf("%s speaker is now OFF\n", s.brand)
}

func (s SmartSpeaker) PlayMusic(song string) {
	fmt.Printf("%s speaker is playing: %s\n", s.brand, song)
}

func main() {
	// Create instances of SmartLight and SmartSpeaker
	light := SmartLight{name: "Living Room"}
	speaker := SmartSpeaker{brand: "JBL"}

	// Use the Light interface for controlling SmartLight
	var l Light = light
	l.TurnOn()
	l.ChangeBrightness(75)
	l.TurnOff()

	fmt.Println()

	// Use the Speaker interface for controlling SmartSpeaker
	var s Speaker = speaker
	s.TurnOn()
	s.PlayMusic("Shape of You")
	s.TurnOff()
}
