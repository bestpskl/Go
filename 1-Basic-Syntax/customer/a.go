/* 	--------------- Package --------------- */

package customer

// 	rules
// 	1. declare with PascalCase (UpperCase at first character) to expose outside package
// 	2. package name is significant but files' name in package are not (they only share resources)
// 	3. use dot syntax to use other packages (packages will import automatically)

var Name = "Customer"           // can use in other files
var firstName = "Customer ABCD" // use in package only

func Multi(a, b int) int {
	return a * b
}

/* 	--------------- Struct & Method --------------- */

// 	getter setter

type Person struct {
	name string // private (not expose)
	age  int    // private (not expose)
}

func (p Person) GetName() string {
	return p.name
}

func (p Person) GetAge() int {
	return p.age
}

func (p *Person) SetName(name string) { // to change value -> not forget to use pointer
	p.name = name
}

func (p *Person) SetAge(age int) { // to change value -> not forget to use pointer
	p.age = age
}
