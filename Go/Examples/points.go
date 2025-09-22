package main

import "fmt"

// Function that accepts a pointer to an integer
func modifyPtr(x *int) {
	*x = 42 //Dereferencing the pointer to change the value
}

func main() {

	x := 10
	fmt.Println("Before Funtion call, x =", x) // 10

	//Pass the addrrss of variable 'a' to the func
	modifyPtr(&x)

	fmt.Println("After Funtion call, x =", x)

	a := 4
	var b float32 = 5
	c := 4.5
	var ptr *int
	country := "India"
	var nation string
	var ptr3 *string = &nation
	var ptr1 *string = &country
	var y int = 45
	var ptr4 *int = &y
	var dblptr **int = &ptr4

	fmt.Println("value of ptr4 =", ptr4)
	fmt.Println("value of dblptr =", dblptr)

	fmt.Println("Valure pointed by ptr1 is:", *ptr1)
	fmt.Println("Address pointed by ptr1 is:", ptr1)
	fmt.Println("Valure pointed by ptr3 is:", *ptr3)
	fmt.Println("Address pointed by ptr3 is:", ptr3)

	var ptr2 *int
	fmt.Println("Value of ptr2 is:", ptr2)

	if ptr2 == nil {
		fmt.Println("ptr2 is null pointer")
	} else {
		fmt.Println("ptr2 is not a null pointer")
	}

	fmt.Printf("type of variable a : %T\n", a)
	fmt.Printf("type of variable b : %T\n", b)
	fmt.Printf("type of variable c : %T\n", c)

	ptr = &a
	fmt.Printf("address of variable a : %p\n", ptr)
	fmt.Printf("Value of a : %d\n", a)
	fmt.Printf("Value of ptr : %d\n", *ptr)

	z := 99
	ptr5 := &z
	ptr6 := &ptr5

	fmt.Println("Value of z :", z)
	fmt.Println("Value pointed by ptr5 :", *ptr5)
	fmt.Println("Address of z :", &z)
	fmt.Println("Value of ptr5 :", ptr5)
	fmt.Println("Address of ptr5 :", &ptr5)
	fmt.Println("Value pointed by ptr6 :", *ptr6)
	fmt.Println("Value pointed by ptr5 via ptr6 :", **ptr6)
}
