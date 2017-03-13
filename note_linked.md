


## 链表
###删除重复的node
dummy node.当链表的头不确定的时候，需要在head之前指定一个新的节点.
```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head==nil{
        return head
    }
    dummyNode:=&ListNode{}
    dummyNode.Next=head
    head=dummyNode
    
    for head.Next!=nil&&head.Next.Next!=nil{
       if(head.Next.Val==head.Next.Next.Val){
           val:=head.Next.Val
           for head.Next!=nil&&head.Next.Val==val{
               head.Next=head.Next.Next
           }
       
           
       }else{
           head =head.Next
       }
        
        
    }
    
    return dummyNode.Next
}

```

1.删除
A.next==空
两个问题：1找到重复的点，删除所有的点。

2.访问.next .val一定要保证该点是不为空的

###反转链表
```
/**
 * URL:https://leetcode.com/problems/reverse-linked-list/?tab=Description
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head==nil|| head.Next==nil{
        return head
    }
    
    var prev *ListNode
    
    for head!=nil{
        temp:=head.Next
        head.Next=prev
        prev=head
        head=temp
    }
    
    return prev //不是head

}

```

t=a
a=b
b=a

###反转链表部分
url:https://leetcode.com/problems/reverse-linked-list-ii/?tab=Description

分为三个部分：
1.找到第一个头，保留尾巴
2.翻转mNode->nNode.
3.接起来
```
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head==nil||m>=n{
        return head
    }
    
    dummy:=&ListNode{}
    dummy.Next=head
    head=dummy  //一定要加上.
    
    for i:=1;i<m;i++{
        if head==nil{
             return nil
        }
        head =head.Next
        
    }
    
   premNode:=head
   mNode:=head.Next
   nNode,postnNode:=mNode,mNode.Next
   for i:=m;i<n;i++{
       if postnNode==nil{
           return nil
       }
       
       temp := postnNode.Next
       postnNode.Next=nNode
       nNode=postnNode
       postnNode=temp
       
   }
   mNode.Next=postnNode
   premNode.Next=nNode
    
    
    return dummy.Next
}

```

### 链表划分

Url:https://leetcode.com/problems/partition-list/?tab=Description
划分两个前后指针即可
注意的地方是需要指定尾巴为nil
```
func partition(head *ListNode, x int) *ListNode {
    if head==nil{
        return head
    }
    less,greater:=&ListNode{},&ListNode{}
    lessNode:=less
    greaterNode:=greater
   for head!=nil{
       if head.Val<x{
           lessNode.Next=head
           lessNode=lessNode.Next
       }else{
           greaterNode.Next=head
           greaterNode=greaterNode.Next
       }
       head=head.Next
   }
   greaterNode.Next=nil //少了这一句要死掉的。
   lessNode.Next=greater.Next
   return less.Next
}
```

### basic skills
1.插入节点
2.删除节点
3.反转链表
4.合并两个链表
5.找到中位数节点
  快慢指针

### dummynode场景：
1. Remove Duplicates from Sorted List II
2. Reverse Linked List II
3. Merge Two Sorted Lists
4. Partition List

1.访问x.next,保证x!=nil
2.删除一个节点
必须知道该节点的前节点是谁
prev-?next-?next
prev->next=prev->next->next
3.如果你的表头会发生变化，那么就用一个dummy node指向原来节点
t=a
a=b
b=t
4.simple/sigl list vs doubly linked list

### sort list
1.找中点
2.sort
3.merge

```
func sortList(head *ListNode) *ListNode {
      if head==nil||head.Next==nil{   //1.head.Next==nil也要加上。
          return head
      }
     
      middle:=findMiddle(head)
      //divide
      right:=sortList(middle.Next) //2.middle.Next不是middle. 递归排序左边 
      middle.Next=nil
      left:=sortList(head)   //递归排序右边
      
      //conquer
      return merge(right,left)  //merge
      
}

func merge(right,left *ListNode) *ListNode{
    dummy:=&ListNode{}
    node:=dummy
    for right!=nil&&left!=nil{
        if right.Val<left.Val{
            node.Next=right
            right=right.Next
        }else{
            node.Next=left
            left=left.Next
        }
        node=node.Next //3.这个记得加上啊，差点跪了。
    }
    if right!=nil{
        node.Next=right
    }
    if left!=nil{
        node.Next=left
    }
    return dummy.Next
}

func findMiddle(head *ListNode) *ListNode{
      slow:=head
      fast:=head.Next
      for(fast!=nil&&fast.Next!=nil){
          fast=fast.Next.Next
          slow=slow.Next
      }
      
      return slow
}
```


### reorder
1.找中点
2.reverse
3.merge


