/*
https://leetcode.com/problems/balanced-binary-tree/?tab=Description

*/
import "math"

func isBalanced(root *TreeNode) bool {
	return maxDepth(root) > -1
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0 //错误点，不是-1，是0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left == -1 || right == -1 || int(math.Abs(float64(left-right))) > 1 {
		return -1
	}

	return int(math.Max(float64(left), float64(float64(right)))) + 1

}