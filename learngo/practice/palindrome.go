package practice

import (
	"strconv"
	"strings"
)

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	//str := fmt.Sprintf("%v", x)
	str := strconv.Itoa(x)
	parts := strings.Split(str, "")
	for i, j := 0, len(parts)-1; i <= j; i, j = i+1, j-1 {
		if parts[i] != parts[j] {
			return false
		}
	}
	return true
}
