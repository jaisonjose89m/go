package main

import (
	"fmt"
	"github.com/jaisonjose89m/go/learngo/practice"
)

func main() {
	invokerAddBinary()
	//invokeSearchInsertPos()
}

func invokerAddBinary() {
	fmt.Println("practice.AddBinary(\"1101\", \"111011\")=", practice.AddBinary("1101", "111011c"))
}

func invokeSearchInsertPos() {
	nums := []int{1, 3, 4, 5}
	fmt.Println("practice.SearchInsert(nums, 2)=", practice.SearchInsert(nums, 2))
}

func invokePalindrome() {
	//fmt.Println("practice.IsPalindrome(121) = ", practice.IsPalindrome(121))
	//fmt.Println("practice.IsPalindrome(-121) = ", practice.IsPalindrome(-121))
	//fmt.Println("practice.IsPalindrome(0) = ", practice.IsPalindrome(0))
	//fmt.Println("practice.IsPalindrome(123) = ", practice.IsPalindrome(123))
}
