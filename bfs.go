BFS有三种方法
1.2个queue
2.1 queue+dummy node
A#B C#D E
3.1 queue(best)

/**
 * url:https://leetcode.com/problems/binary-tree-level-order-traversal/?tab=Description
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {

    result:=[][]int{}
    if root==nil{
        return result
    }
    size:=0
    q:=make(chan *TreeNode,1000)
    q<-root
    size++
    for size>0{
        level:=[]int{}
        length:=size
        for i:=0;i<length;i++{
               v:=<-q
               level=append(level,v.Val)
               if v.Left!=nil{
                   q<-v.Left
                   size++
               }
               if v.Right!=nil{
                   q<-v.Right
                   size++
               }

           }
           size=size-length
           result=append(result,level)
        }


     return result

}



