package magazine

type Subscriber struct {
	Name string
	Rate float64
	Active bool
	HomeAddress Address
}
type Employee struct {
	Name string
	Salary float64
	HomeAddress Address
}

type Address struct {
	Street string
	City string
	State string
	PostalCode string
}

/*Go type names follow the same rule as variable and function names: if the name
of a variable, function, or type begins with a capital letter, it is considered
exported and can be accessed from outside the package it’s declared in.*/

/*Even if a struct type is exported from a package, its fields will be unexported if
their names don’t begin with a capital letter.*/

/*Struct field names must also be capitalized if you want to export them from
their package.*/

//A defined type’s name must be capitalized to be exported
//Struct field names must be capitalized to be exported