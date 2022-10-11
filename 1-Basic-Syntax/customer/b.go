package customer

// 	rules
// 	1. declare with PascalCase (UpperCase at first character) to expose outside package
// 	2. package name is significant but files' name in package are not (they only share resources)
// 	3. use dot syntax to use other packages (packages will import automatically)

func Hello() string {
	return "Hello " + firstName
}
