package primitives

import (
	"fmt"
)

func Primitive() {
	fmt.Println("\n*** Output from Primitive ***")
	// classic declaration
	var integer int
	integer = 7
	fmt.Println(integer)
	// classic declaration with initialization
	var pi float32 = 3.14
	fmt.Println(pi)

	// shorthand go declaration
	firstName := "Jaison"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Print(c)

	r, i := real(c), imag(c)
	fmt.Print(" Real: ", r)
	fmt.Println(" Imaginary: ", i)
}
