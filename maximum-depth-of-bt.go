/*
分治
https://leetcode.com/problems/maximum-depth-of-binary-tree/?tab=Description
*/import "math"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return 1 + int(math.Max(float64(left), float64(right)))

}


