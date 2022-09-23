package practice

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPathSum(root *TreeNode, targetSum int) bool {
	stack := make([]int, 0)
	stack = append(stack, 1)
	fmt.Println(stack)
	return false
}
