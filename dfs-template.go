
模板1:遍历

func traverse(root *TreeNode){
    if root==nil{
        return
    }

    traverse(root.Left)
    traverse(root.Right)
}

模板2：分治

func traversal(root *TreeNode) *ResultType{
    //空 或者 叶子
    if root==nil{
        //做点啥 return
    }
    //分
    ResultType left=traversal(root.Left)
    ResultType right = traversal(root.Right)

    //治
    ResultType result = Merge from left & right
    return ResultType
}



