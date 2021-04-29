package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactDetails
}

type contactDetails struct {
	email   string
	zipCode int
}

func main() {
	//alex := person{firstName :"Alex",lastName: "Anderson"}
	var alex person
	fmt.Printf("%+v", alex)
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "abc@gmail.com"
	alex.contact.zipCode = 544333

	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contact: contactDetails{
			email:   "abc@gmail.com",
			zipCode: 454545,
		},
	}
	//jimPointer := &jim //address at which jim is stored
	jim.updateName("jimmy")
	jim.print()
	alex.print()
}

//pass by value
// func (p person) updateName(newFirstName string){
// 	p.firstName = newFirstName
// }

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p.firstName, p.lastName)
	fmt.Printf("%+v", p)
}
