/**
 * Definition for undirected graph.
 * class UndirectedGraphNode {
 *     int label;
 *     List<UndirectedGraphNode> neighbors;
 *     UndirectedGraphNode(int x) { label = x; neighbors = new ArrayList<UndirectedGraphNode>(); }
 * };
 */
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

public class Solution {
    
  //Definition for singly-linked list with a random pointer.
  class RandomListNode {
      int label;
      RandomListNode next, random;
      RandomListNode(int x) { this.label = x; }
  };
 
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

   
  
    
 }   