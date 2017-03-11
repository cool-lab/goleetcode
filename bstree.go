package leetcode

/*
https://leetcode.com/problems/validate-binary-search-tree/
校验正确性
*/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	list := []int{}
	inOrderTraversal(root, &list)
	for i := 1; i < len(list); i++ {
		if list[i] <= list[i-1] {
			return false
		}
	}
	return true

}

func inOrderTraversal(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left, list)
	*list = append(*list, root.Val)
	inOrderTraversal(root.Right, list)
}
