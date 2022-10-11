package main // always

import ( // import automatically

	"basicsyntax/customer"
	"fmt"
	"unicode/utf8"
)

type Person struct {
	Name string
	Age  int
	Pass []string
}

func (p Person) Hello() string {
	return "Hello " + p.Name
}

func main() {

	// 	By Pasakorn Limchuchua

	//	Main Reference
	// 	-> https://github.com/codebangkok/golang
	// 	-> https://www.youtube.com/watch?v=JbIS97exQnQ&ab_channel=CodeBangkok
	// 	-> https://go.dev/ref/spec
	//	-> https://pkg.go.dev/builtin
	// 	-> https://pkg.go.dev/fmt
	//	-> https://pkg.go.dev/unicode/utf8

	//	Extensions
	// 	-> Go, Error Lens

	// 	Open / Close terminal 			-> ctrl + `
	// 	Help Command 					-> ctrl + spacebar

	// 	Commands
	// 	go mod init [module_name]		-> create modules
	// 	go run [file_name].go			-> run file
	//	go run .						-> run file

	/* 	--------------- Output ---------------	*/

	// 	package fmt (ref)

	// 	%v 		-> for value in default format
	// 	%#v 	-> for deep %v
	// 	%T 		-> for type of value

	fmt.Printf("Hello %v\n", 10) // Hello 10
	fmt.Println("Hello world")   // Hello world

	/* 	--------------- Type, Variable Declaration & Operators --------------- */

	// 	bool 										(zero value : false)
	// 	string 										(zero value : "")
	// 	int int8 int16 int32(rune) int64 (+-0) 		(zero value : 0)
	// 	uint uint8(byte) uint16 uint32 uint64 (+) 	(zero value : 0)
	// 	float32 float64 							(zero value : 0)
	// 	complex64 complex128 						(zero value : 0 + 0i)

	// 	var		-> can change value
	// 	const	-> can't change value

	// 	standard declaration 	-> use anywhere
	var x int = 10
	_ = x // -> to ignore compile time when not used
	//	type of influence
	var x2 = 10
	_ = x2

	// 	short declaration 		-> use only in func
	y := 10
	fmt.Println(y) // 10

	// 	Arithmetic operators	-> apply to numeric values and yield a result of the same type as the first operand
	// 	+    sum                    integers, floats, complex values, strings
	// 	-    difference             integers, floats, complex values
	// 	*    product                integers, floats, complex values
	// 	/    quotient               integers, floats, complex values
	// 	%    remainder              integers

	//	Comparison operators	-> compare two operands and yield an untyped boolean value
	// 	==    equal
	// 	!=    not equal
	// 	<     less
	// 	<=    less or equal
	// 	>     greater
	// 	>=    greater or equal

	//	Logical operators		-> apply to boolean values and yield a result of the same type as the operands
	// 	&&    conditional AND    p && q  is  "if p then q else false"
	// 	||    conditional OR     p || q  is  "if p then true else q"
	// 	!     NOT                !p      is  "not p"

	/* 	--------------- Condition --------------- */

	// 	if, else if, else
	point := 50
	if point >= 50 && point <= 100 {
		fmt.Println("Pass") // Pass
	} else if point >= 20 {
		fmt.Println("Good")
	} else {
		fmt.Println("Try again")
	}

	// 	switch
	switch day := "Monday"; day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend")
	case "Friday":
		fmt.Println("Workday")
	default:
		fmt.Println("!!") // !!
	}

	/* 	--------------- Array --------------- */

	array1 := [3]int{}                //
	fmt.Printf("%v\n", array1)        // [0 0 0]
	array2 := [3]int{1, 2, 3}         // initialize value
	fmt.Printf("%#v\n", array2)       // [3]int{0, 0, 0}
	array3 := [...]int{1, 2, 3, 4, 5} // auto count
	array3[4] = 10                    //
	fmt.Printf("%#v\n", array3)       // [5]int{10, 2, 3, 4, 5}

	// 	2d array
	array4 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(array4) // [[1 2 3] [4 5 6] [7 8 9]]

	/* 	--------------- Slice --------------- */

	slice1 := []int{1, 2, 3}
	fmt.Printf("%#v\n", slice1) // []int{1, 2, 3}

	// 	append 			-> appends elements to the end of a slice
	slice1 = append(slice1, 4)  // add element
	fmt.Printf("%#v\n", slice1) // []int{1, 2, 3, 4}

	// 	len 				-> returns the length (beware unicode)
	length := len(slice1)
	fmt.Printf("%#v, len = %v\n", slice1, length) // []int{1, 2, 3, 4}, len = 4

	// 	package utf8 (ref)

	name := "ภาสกร"                           // beware unicode !!
	fmt.Println(len(name))                    // 15
	fmt.Println(utf8.RuneCountInString(name)) // 5

	//				 0	 1	 2	 3	 4	 5	 6	 7	 8
	slice2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	slice2_1 := slice2[1:]
	fmt.Println(slice2_1) // [20 30 40 50 60 70 80 90]
	slice2_2 := slice2[1:4]
	fmt.Println(slice2_2) // [20 30 40]
	slice2_3 := slice2[:6]
	fmt.Println(slice2_3) // [10 20 30 40 50 60]
	slice2_4 := slice2[:]
	fmt.Println(slice2_4) // [10 20 30 40 50 60 70 80 90]

	// 	remove
	slice3 := append(slice2[:4], slice2[5:]...)
	fmt.Println(slice3) // [10 20 30 40 60 70 80 90]

	/* 	--------------- Map --------------- */

	countries := map[string]string{}
	countries["th"] = "Thailand"
	countries["en"] = "United State"
	fmt.Println(countries["th"]) // Thailand

	// 	a, b := 10, 20
	country, ok := countries["jp"] // return 2 values
	if ok {
		fmt.Println(country)
	} else {
		fmt.Println("no value") // no value
	}

	/* 	--------------- For Loop --------------- */

	// 	for
	slice4 := []int{10, 20, 30, 40, 50}
	for i := 0; i < len(slice4); i++ {
		fmt.Println(slice4[i])
		// 10
		// 20
		// 30
		// 40
		// 50
	}

	// 	while
	i := 0
	for i < len(slice4) {
		fmt.Println(slice4[i])
		i++
		// 10
		// 20
		// 30
		// 40
		// 50
	}

	// 	for each
	for i, v := range slice4 {
		fmt.Println(i, v)
		// 0 10
		// 1 20
		// 2 30
		// 3 40
		// 4 50
	}
	for _, v := range slice4 {
		fmt.Println(i, v)
		// 10
		// 20
		// 30
		// 40
		// 50
	}

	/* 	--------------- Function --------------- */

	// 	void
	/*
		func sum1() int {
			a1 := 10
			b1 := 5
			return a1 + b1
		}
	*/
	c1 := sum1()
	fmt.Println(c1) // 15

	// 	parameter & return
	/*
		func sum2(a2 int, b2 int) int {
			return a2 + b2
		}
	*/
	c2 := sum2(10, 5)
	fmt.Println(c2) // 15

	/*
		func sum3(a3, b3 int) (int, string, bool) {
			return a3 + b3, "Hello", true
		}
	*/
	c3, d3, e3 := sum3(10, 5)
	fmt.Println(c3, d3, e3) // 15 Hello true

	// 	anonymous function		-> function as variable (no name)
	f1 := func(a, b int) int {
		return a + b
	}
	sum4 := f1(10, 20)
	fmt.Println(sum4) // 30

	// 	closure function 		-> function in function (as parameter + assume function is one type)
	/*
		// (int,int)int
		func add(aa, bb int) int {
			return aa + bb
		}

		// (int,int)int
		func sub(aa, bb int) int {
			return aa - bb
		}

		func cal(f2 func(int, int) int) { 	// f2 is name of func(int, int) int
			result := f2(50, 10)
			fmt.Println(result)
		}

		*** note that 'cal' will receive only type of (int, int) int) -> add & sub are ok***
	*/
	cal(add) // 60
	cal(sub) // 40

	// 	closure x anonymous
	/*
		func cal(f2 func(int, int) int) { 	// f2 is name of func(int, int) int
			result := f2(50, 10)
			fmt.Println(result)
		}
	*/
	f3 := func(a, b int) int {
		return (a + a + b + b)
	}

	cal(f3) // 120

	/*
		func cal(f2 func(int, int) int) { 	// f2 is name of func(int, int) int
			result := f2(50, 10)
			fmt.Println(result)
		}
	*/

	cal(func(a, b int) int { // can omit f3
		return -a - a - b - b
	}) // 120

	// 	variadic function		-> function with unlimited parameters
	/*
		func sumElement1(a []int) int {
			s := 0
			for _, v := range a {
				s += v
			}
			return s
		}
	*/
	sliceForSum := []int{100, 200, 300}
	ans1 := sumElement1(sliceForSum)
	fmt.Println(ans1) // 600

	/*
		func sumElement2(a ...int) int {
			s := 0
			for _, v := range a {
				s += v
			}
			return s
		}
	*/
	ans2 := sumElement2(1000, 2000, 3000, 4000, 5000) // do not need to declare slice/array first
	fmt.Println(ans2)                                 // 15000

	/* 	--------------- Package --------------- */

	// 	rules
	// 	1. declare with PascalCase (UpperCase at first character) to expose outside package
	// 	2. package name is significant but files' name in package are not (they only share resources)
	// 	3. use dot syntax to use other packages (packages will import automatically)

	// 	*** note that 'package customer' includes a.go b.go ***

	fmt.Println(customer.Name)    // Customer
	fmt.Println(customer.Hello()) // Hello Customer ABCD
	ans3 := customer.Multi(5, 5)  //
	fmt.Println(ans3)             // 25

	/* 	--------------- Pointer --------------- */

	// 	*** value type *** -> copy with (=) will copy value 'not' memory address

	// 	bool 										(zero value : false)
	// 	string 										(zero value : "")
	// 	int int8 int16 int32(rune) int64 (+-0) 		(zero value : 0)
	// 	uint uint8(byte) uint16 uint32 uint64 (+) 	(zero value : 0)
	// 	float32 float64 							(zero value : 0)
	// 	complex64 complex128 						(zero value : 0 + 0i)
	// 	struct

	// 	*** reference type *** -> copy with (=) will copy memory address

	//  array
	//	slice
	//	map
	// 	function
	// 	pointer

	// 	to show memory address 	-> use & before var name
	// 	to show value 			-> use * before pointer name

	var var1, var2 int
	var1 = 10
	var2 = var1
	fmt.Println(var1)  // 10 				-> (value)
	fmt.Println(var2)  // 10 				-> (value)
	fmt.Println(&var1) // 0xc0000123c0 		-> (memory address)
	fmt.Println(&var2) // 0xc0000123c8 		-> (memory address)
	// &var1 &var2 have different memory address
	var1 = 20
	fmt.Println(var1)  // 20
	fmt.Println(var2)  // 10
	fmt.Println(&var1) // 0xc0000123c0
	fmt.Println(&var2) // 0xc0000123c8
	// change value of var1 will 'not' effect value of var2 because they have different memory address

	var var3 int = 10
	var var4 *int = &var3
	fmt.Println(var3)  // 10				-> (value)
	fmt.Println(*var4) // 10				-> (value)
	fmt.Println(&var3) // 0xc0000ba3c0		-> (memory address)
	fmt.Println(var4)  // 0xc0000ba3c0		-> (memory address)
	// &var3 var4 have same memory address
	*var4 = 20
	fmt.Println(var3)  // 20
	fmt.Println(*var4) // 20
	fmt.Println(&var3) // 0xc0000ba3c0
	fmt.Println(var4)  // 0xc0000ba3c0
	// change value of var4 will effect value of var3 because they have same memory address

	/*
		func sumPointer(result *int) { // -> receive parameter with pointer
			a := 10
			b := 20
			*result = a + b // -> must not use return (value will change automatically through memory address)
		}
	*/

	var var5 int = 0
	sumPointer(&var5) // send memory address
	fmt.Println(var5) // 30

	/* 	--------------- Struct & Method --------------- */

	// 	struct -> use struct in go instead of class in other languages

	// 	rules
	//	1. has only field
	//	2. declare with PascalCase (UpperCase at first character) to expose outside (public)
	//	3. use dot syntax to access fields

	/*
		type Person struct { 	// public
			Name string			// public
			Age  int			// public
			Pass []string		// public
		}
	*/

	struct1 := Person{
		Name: "Best",
		Age:  22,
		Pass: []string{"en", "th"},
	}
	fmt.Println(struct1) // {Best 22 [en th]}

	struct2 := struct1
	struct1.Name = "Noon"
	struct1.Pass = []string{"th", "en"}
	fmt.Println(struct1.Name) // Noon
	fmt.Println(struct2.Name) // Best
	fmt.Println(struct1.Pass) // [th en]
	fmt.Println(struct2.Pass) // [en th]
	// 	*** Note that struct is value type ***

	// 	method (receiver function) -> extension to struct as function

	/*
		type Person struct {
			Name string
			Age  int
			Pass []string
		}

		func (p Person) Hello() string {	// -> attach struct (receiver) to function in front of function name
			return "Hello " + p.Name
		}
	*/
	fmt.Println(struct1.Hello()) // Hello Noon

	//	getter setter
	struct3 := customer.Person{}
	struct3.SetName("Pasakorn")
	fmt.Println(struct3.GetName()) // Pasakorn
}

func sum1() int {
	a1 := 10
	b1 := 5
	return a1 + b1
}

func sum2(a2 int, b2 int) int {
	return a2 + b2
}

func sum3(a3, b3 int) (int, string, bool) {
	return a3 + b3, "Hello", true
}

// (int,int)int
func add(aa, bb int) int {
	return aa + bb
}

// (int,int)int
func sub(aa, bb int) int {
	return aa - bb
}

func cal(f2 func(int, int) int) {
	result := f2(50, 10)
	fmt.Println(result)
}

func sumElement1(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func sumElement2(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func sumPointer(result *int) {
	a := 10
	b := 20
	*result = a + b
}
