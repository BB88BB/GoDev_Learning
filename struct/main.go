package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {

	// Initialize and assign value
	/*
		alex := person{firstname: "Alex", lastname: "Anderson"}

		var tom, vivian person
		tom.firstname = "Tom"
		tom.lastname = "Anderson"

		fmt.Println(alex)
		// {Alex Anderson}

		fmt.Printf("%+v\n", alex)
		// {firstname:Alex lastname:Anderson}%

		fmt.Printf("%+v\n", vivian)
		// {firstname: lastname:}%

		fmt.Printf("%#v", tom)
		// main.person{firstname:"Tom", lastname:"Anderson"}%
	*/

	// Embedded struct
	jim := person{
		firstname: "Jim",
		lastname:  "Andrew",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.print()

	// Pointer
	/*
		jimPointer := &jim
		jimPointer.updateName("jimmy")
		jim.print()
	*/

	// Pointer shortcut, will allow to call receiver *person
	jim.updateName("jimmy")
	jim.print()
}

// ptrToPerson *person -> pointer of person type can call this receiver, BTW shortcut also can
func (ptrToPerson *person) updateName(newfirstname string) {

	// get the value of this pointer
	(*ptrToPerson).firstname = newfirstname
}

func (p person) print() {

	fmt.Printf("%+v", p)
}
