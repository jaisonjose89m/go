package flow

import "fmt"

func SimpleFor() {
	fmt.Println("From SimpleFor")
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
		if i == 3 {
			// break
			continue
		}
		fmt.Println("completed loop body for ", i)
	}
}

func ClassicFor() {
	fmt.Println("From ClassicFor")
	for i := 1; i < 5; i++ {
		fmt.Println(i)
		i++
		if ((i == 3) || (i == 4)) && i < 10 {
			// break
			continue
		} else {
			fmt.Println("Condition doesn't match")
		}
		fmt.Println("completed loop body for ", i)
	}
}

func InfiniteLoop() {
	fmt.Println("From InfiniteLoop")
	var i int
	for {
		if i > 4 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func LoopOverSlice() {
	fmt.Println("From LoopOverSlice")
	slice := []int{1, 2, 3}
	for i, v := range slice {
		fmt.Println(i, v)
	}
}

func LoopOverMap() {
	sampleMap := map[string]int{"http": 8080, "ftp": 80}
	for k, v := range sampleMap {
		fmt.Println(k, v)
	}
}
