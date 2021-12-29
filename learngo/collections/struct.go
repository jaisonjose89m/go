package learngo

import "fmt"

func Struct() {
	fmt.Println("\n*** Output from Struct ***")
	type user struct {
		ID         int
		FirstName  string
		SecondName string
	}
	var u1 user
	// classic assignment
	u1.ID = 1
	u1.FirstName = "Samuel"
	u1.SecondName = "John"
	fmt.Println(u1)
	// one line shorthand
	u2 := user{ID: 1, FirstName: "Samuel", SecondName: "John"}
	fmt.Println(u2)

	// multi line shorthand
	u3 := user{
		ID:         1,
		FirstName:  "Samuel",
		SecondName: "John",
	}
	fmt.Println(u3)
}
