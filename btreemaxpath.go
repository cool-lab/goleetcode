/*
处理叶子，处理null,处理null更好.★
首先，想一个简化版(single path)，找从root到任意点得最大值。类似于maxDepth，每次加root.val而不再是+1
求单路的时候，如果root加左儿子单路或者右儿子单路最后的值都小于0，则返回0，意味着不要root开始的这个单路了
本题思路，divide & conquer
求最大路径和就等于下面三个值的最大值：
左子树的最大路径和
右子树最大路径和
左子树单路 + 右子树单路 + root.val
*/
import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type ResultType struct {
	SinglePath int //通过root节点的最长路径
	MaxPath    int
}

func NewResultType(single, max int) *ResultType {
	return &ResultType{
		SinglePath: single,
		MaxPath:    max,
	}
}

func maxPathSum(root *TreeNode) int {
	return helper(root).MaxPath
}

func helper(root *TreeNode) *ResultType {
	if root == nil {
		return NewResultType(0, -100000000)
	}
	// Divide
	left := helper(root.Left)
	right := helper(root.Right)
	singlePathSum := int(math.Max(float64(left.SinglePath), float64(right.SinglePath))) + root.Val
	singlePathSum = int(math.Max(float64(singlePathSum), 0.0))
	maxPath := int(math.Max(float64(left.MaxPath), float64(right.MaxPath)))                       //左右最大的单个最长路径
	maxPath = int(math.Max(float64(maxPath), float64(left.SinglePath+right.SinglePath+root.Val))) //2.比较最大maxpath.与singlepath

	return NewResultType(singlePathSum, maxPath)

}