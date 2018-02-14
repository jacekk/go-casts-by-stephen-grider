package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}

func main() {
	/*
		type simplePerson struct {
			firstName string
			lastName  string
		}
		// alex := simplePerson{"John", "Doe"} // not recommended; order depended
		john := simplePerson{firstName: "John", lastName: "Doe"}
		var jane simplePerson

		fmt.Println(john)
		fmt.Println(jane)
		fmt.Printf("%+v\n", john)

		john.firstName = "JohnTheSecond"
		fmt.Printf("%+v\n", john)
	*/

	jim := person{
		firstName: "Jim",
		lastName:  "Carrey",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	jimPointer := &jim
	jimCopy := jim

	jimPointer.updateName("Jimmy")

	jim.print()        // Jimmy
	jimPointer.print() // Jimmy
	jimCopy.print()    // Jim
}

/*
Noted:

Value types:
	- int
	- float
	- string
	- bool
	- struct

Reference types:
	- slices
	- maps
	- channles
	- pointers
	- functions

*/
