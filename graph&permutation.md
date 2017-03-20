# 图搜索

## Clone Graph

由于没有对go的支持，所以改为java版本

```java
public class Solution{
    /*
    */

    public UndirectedGraphNode cloneGraph(UndirectedGraphNode node){
        if (node==null){
            return node;
        }

        //使用 bfs 算法遍历获得所有的node.
        ArrayList<UndirectedGraphNode> nodes =getNodes(node);

        // copy nodes, 存储老的-》新的hash map 
        HashMap<UndirectedGraphNode,UndirectedGraphNode> mapping =new HashMap<>();
        for(UndirectedGraphNode n:nodes){
            mapping.put(n,new UndirectedGraphNode(n.label));
        }

        // 复制edges
        for(UndirectedGraphNode n:nodes){
            UndirectedGraphNode newNode =mapping.get(n);
            for(UndirectedGraphNode neighbor:n.neighbors){
                UndirectedGraphNode newNeighbor = mapping.get(neighbor);
                newNode.neighbors.add(newNeighbor);
            }
        }
        return mapping.get(node);
    }
    
    private ArrayList<UndirectedGraphNode> getNodes(UndirectedGraphNode node){
        Queue<UndirectedGraphNode> queue =new LinkedList<UndirectedGraphNode>();
        HashSet<UndirectedGraphNode> set =new HashSet<>();
        
        queue.offer(node);
        set.add(node);
        while (!queue.isEmpty()) {
            UndirectedGraphNode head =queue.poll();
            for(UndirectedGraphNode neighbor:head.neighbors){
                if(!set.contains(neighbor)){
                    set.add(neighbor);
                    queue.offer(neighbor);
                }
            }
        }

        return new ArrayList<UndirectedGraphNode>(set);
    }
}
```

## clone random pointer

```java
/**
 * Definition for singly-linked list with a random pointer.
 * class RandomListNode {
 *     int label;
 *     RandomListNode next, random;
 *     RandomListNode(int x) { this.label = x; }
 * };
 */
public class Solution {
    public RandomListNode copyRandomList(RandomListNode head) {
        if (head==null){
            return head; 
        }
        HashMap<RandomListNode,RandomListNode> mapping = new HashMap<RandomListNode,RandomListNode>();
        RandomListNode node=head;
        
        while(node!=null){
            RandomListNode newNode=new RandomListNode(node.label);
            mapping.put(node,newNode);
            node=node.next;
        }

        node = head;
        
        while(node!=null){
            RandomListNode  newNode=mapping.get(node);
            newNode.next=mapping.get(node.next);
            newNode.random=mapping.get(node.random);
            node=node.next;
        }

        return mapping.get(head);

    }
    
  
    
 }   


```

### BFS

Idea: 从某个点出发， 找到其他所有的点。
Compare: Graph bfs vs Tree bfs.

## 全排列
全排列，一定有一个for循环，一个一个放入。

```
package permutations

func permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	res := [][]int{}
	list := []int{}
	dfs(&res, list, nums)
	return res
}

func dfs(res *[][]int, list, nums []int) {
	if len(list) == len(nums) {
		replia := make([]int, len(list)) //要进行深度copy.不然会重复。
		copy(replia, list)
		*res = append(*res, replia)
		return
	}
	for i := 0; i < len(nums); i++ {
		if contains(nums[i], list) {        //如果已经有了，则不需要再加入了.
			continue
		}
		list = append(list, nums[i])    //加入列表
		dfs(res, list, nums)
		list = list[0 : len(list)-1]

	}
	return
}

func contains(num int, nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}


```

## subsets

```
package subsets

func subsets(nums []int) [][]int {
	subset := []int{}
	return dfs(nums, subset, 0)
}

func dfs(nums, subset []int, pos int) (ret [][]int) {

	setCopy := make([]int, len(subset))
	copy(setCopy, subset)
	ret = append(ret, setCopy)

	for i := pos; i < len(nums); i++ {
		subset = append(subset, nums[i])
		ret = append(ret, dfs(nums, subset, i+1)...) //注意返回需要append到内容中
		subset = subset[:len(subset)-1]
	}
	return
}


```

### permutation子问题之八皇后问题

```
package nqueen

func solveNQueens(n int) [][]string {

	list := []int{}

	return dfs(list, n)
}

func dfs(list []int, n int) (result [][]string) {
	if len(list) == n {
		result = append(result, (translate(list, n)))
	}
	for col := 0; col < n; col++ {
		if !isValid(list, col) {
			continue
		}
		list = append(list, col)
		result = append(result, dfs(list, n)...)
		list = list[:len(list)-1]
	}
	return
}

func translate(list []int, n int) (ret []string) {

	for _, value := range list {
		b := make([]byte, n)
		for j := 0; j < n; j++ {
			if j != value {
				b[j] = '.'
			} else {
				b[j] = 'Q'
			}
		}

		ret = append(ret, string(b))
	}
	return
}

func isValid(list []int, col int) bool {
	row := len(list)
	for i := 0; i < row; i++ {
		if (list[i] == col) || (i-list[i] == row-col) || (i+list[i] == row+col) {
			return false
		}
	}
	return true
}


```

## subsets2问题

```
package subsets2

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {

	sort.Ints(nums)
	return dfs(nums, []int{}, 0)
}

func dfs(nums, list []int, pos int) (result [][]int) {
	copyList := make([]int, len(list))
	copy(copyList, list)
	result = append(result, copyList)

	for i := pos; i < len(nums); i++ {
		if i != pos && nums[i] == nums[i-1] { //当不是第一元素，且是第一个元素之后的相识元素都不放入。
			continue
		}
		list = append(list, nums[i])
		result = append(result, dfs(nums, list, i+1)...)
		list = list[:len(list)-1]
	}

	return result
}


```


## 总结套路
permutation问题是dfs(nums,list,n)
其中，nums是总的列表，list为每次加入到排序的列表，满足到list长度等于nums长度才可以保存，n为恒定值. n!,复杂度
subset 每次dfs(nums,list,pos)每次都要保存子集，pos每次递增。
permutation套路：

```
func dfs(res *[][]int, list, nums []int) {
	if len(list) == len(nums) {
		replia := make([]int, len(list)) //要进行深度copy.不然会重复。
		copy(replia, list)
		*res = append(*res, replia)
		return
	}
	for i := 0; i < len(nums); i++ {
		if contains(nums[i], list) {        //如果已经有了，则不需要再加入了.
			continue
		}
		list = append(list, nums[i])    //加入列表
		dfs(res, list, nums)
		list = list[0 : len(list)-1]   //从列表中删除

	}
	return
}
```

subsets套路：
```
func dfs(nums, subset []int, pos int) (ret [][]int) {

	setCopy := make([]int, len(subset))
	copy(setCopy, subset)
	ret = append(ret, setCopy)

	for i := pos; i < len(nums); i++ {
		subset = append(subset, nums[i])
		ret = append(ret, dfs(nums, subset, i+1)...) //注意返回需要append到内容中
		subset = subset[:len(subset)-1]
	}
	return
}
```






