package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	james := person{
		firstName: "James",
		lastName:  "Erwin",
		contact: contactInfo{
			email:   "email.com",
			zipCode: 1234,
		},
	}
	var michael person
	jamesPointer := &james

	fmt.Println(james, michael)
	//%+v will print out all field names plus their values
	fmt.Printf("%+v", michael)
	fmt.Printf("%+v", james)
	james.print()
	jamesPointer.updateName("Jimmy")
	james.print()

	//Notice that we don't have to pass in a pointer!:
	james.updateName("Alexa")
	james.print()
}

//Need to pass in a pointer to gain reference
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

//Adding a receiver to a struct
func (p person) print() {
	fmt.Printf("%+v", p)
}
