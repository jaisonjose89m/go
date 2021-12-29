package learngo

import "fmt"

func Slice() {
	fmt.Println("\n*** Output from Slice ***")
	// slice with explicit backing array
	arr := [2]int{1, 2}
	s1 := arr[:]
	fmt.Println("arr:", arr)
	fmt.Println("s1:", s1)

	// slice with go managed backing array
	s2 := []int{1, 2, 3}
	fmt.Println("s2:", s2)

	s2 = append(s2, 4, 5)
	fmt.Println("s2:", s2)

	s3 := s2[:2]  // starting from 0 to index 2 but excluding it
	s4 := s2[3:]  // starting from 3 to last index
	s5 := s2[1:4] // starting from 1 to index 4 but excluding it
	fmt.Println("s3:", s3, "s4:", s4, "s5:", s5)

}
