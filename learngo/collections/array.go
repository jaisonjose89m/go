package learngo

import "fmt"

func Array() {
	fmt.Println("\n*** Output from Array ***")
	/*
		Array is fixed length
		Index starting at 0
	*/
	// long form
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println("arr:", arr)

	// short form
	arr2 := [2]int{1, 2}
	fmt.Println("arr2:", arr2)
}
