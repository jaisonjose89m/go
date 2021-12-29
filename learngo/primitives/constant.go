package primitives

import "fmt"

const (
	firstConst = iota
	secondConst
	thridConst
)

const (
	fourthConst = iota
	fifthConst
)

func Constant() {
	fmt.Println("\n*** Output from Constant ***")
	// constant without type
	const pi = 3.14
	fmt.Println(pi + 1)
	fmt.Println(pi + 2.14)

	// constant with type
	const c int = 45
	fmt.Println(c + 14)
	fmt.Println(float32(c) + 12.5)

	// use of iota - 0,1,2
	// resets in every const blocks
	fmt.Println(firstConst, secondConst, thridConst)
	fmt.Println(fourthConst, fifthConst)
}
