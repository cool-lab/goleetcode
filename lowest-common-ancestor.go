/*
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/?tab=Description
*/
func lowestCommonAncestor(root, node1, node2 *TreeNode) *TreeNode {
	if root == nil || root == node1 || root == node2 {
		return root
	}
    // Divide
    left:=lowestCommonAncestor(root.Left,node1,node2）
    right:=lowestCommonAncestor(root.Right,node1,node2)

    //Conquer
    if (left!=nil&&right!=nil){
        return root
    }
    if (left!=nil){
        return left
    }
    if right!=nil{
        return right
    }
    return nil
}