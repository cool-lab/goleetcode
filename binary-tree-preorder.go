/*
先序遍历

*/
func preorderTraversal(root *TreeNode) []int {

	list := []int{}
	preorder(root, &list)
	return list
}

func preorder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}

	*list = append(*list, root.Val)

	if root.Left != nil {
		preorder(root.Left, list)
	}
	if root.Right != nil {
		preorder(root.Right, list)
	}
	return
}

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
    非递归版本，迭代版本。
*/

func preorderTraversal(root *TreeNode) []int {
	list := []int{}
	var node *TreeNode
	if root == nil {
		return list
	}
	queue := []*TreeNode{root}

	for {
		if len(queue) == 0 {
			return list
		}
		node, queue = queue[len(queue)-1], queue[:len(queue)-1] //出 1.注意

		list = append(list, node.Val)
		if node.Right != nil { //先右边，再左边。
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}

	return list

}








