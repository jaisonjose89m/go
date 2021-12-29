package learngo

import "fmt"

func Pointer() {
	// pointer operator *
	fmt.Println("\n*** Output from Poniter ***")
	var firstName *string = new(string)
	*firstName = "Jaison"
	fmt.Print(firstName)
	fmt.Println("", *firstName)
	// address of operartor &
	secondName := "Jose"
	ptr := &secondName
	fmt.Println(ptr, *ptr)

	secondName = "Joseph"
	fmt.Println(ptr, *ptr)

	a := 34
	p := &a
	fmt.Println(a, p, *p)

	var sptr *string
	sptr = &secondName
	fmt.Println(*sptr)
}
