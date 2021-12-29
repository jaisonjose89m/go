package collections

import "fmt"

func Map() {
	fmt.Println("\n*** Output from Map ***")
	m := map[string]int{"foo": 11}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["bar"] = 23
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
}
