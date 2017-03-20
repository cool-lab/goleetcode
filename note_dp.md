
面试中常见动态规划的分类
1.矩阵上的动态规划
2.序列上的动态规划

## triangle
状态方程
fn(i,j)= max(fn(i,j+1)+a[i,j],fn(i+1,j+1)+a[i,j])
```
https://leetcode.com/problems/triangle/#/description
func minimumTotal(triangle [][]int) int {
    
  high:=len(triangle)
  length:=len(triangle[high-1])
  sum:=triangle[length-1]
  for i:=high-2;i>=0;i--{
      for j:=0;j<=i;j++{
          if (sum[j]+triangle[i][j])<(sum[j+1]+triangle[i][j]){
              sum[j]=sum[j]+triangle[i][j]
          }else{
               sum[j]=sum[j+1]+triangle[i][j]
          }
      }
  }
  return sum[0]
}

```

### DFS模板

```
func dfs(x,y,sum int){
    if x==n{
        if sum<best{
            best=sum
        }
     return
    }

    dfs(x+1,y,sum+a[x][y])
    dfs(x+1,y+1,sum+a[x][y])
}

dfs(0,0)

```

### 分治

```
func dfs(x,y) int{
    if(x==n){
        return 0
    }

    return min(dfs(x+1,y),dfs(x+1,y+1))+a[x][y]
}

dfs(0,0)

```

### DFS优化版(记忆化搜索版)

```
// Divide & Conquer
// 从x,y出发走到最底层所能找到的最小路径和

func dfs(x,y int) int{
    if x==n{
        return 0
    }
    if (hash[x][y]!=-1){  //极大提升效率，这个效率和dp差不多了。
        return hash[x][y]
    }

    hash[x][y]=min(dfs(x+1,y),dfs(x+1,y+1))+a[x][y]
    return hash[x][y]
}



```

记忆化搜索
本质上：动态规划

动态规划就是解决了重复计算的搜索
动态规划的实现方式：
1.记忆化搜索
2.循环

####自底向上搜索
```
A[][]
f[i][j]从i,j出发走到最后一层的最小路径
for(int i=0;i<n;i++){
    f[n-1][i]=A[n-1][i];
    //最底先有值
}
//循环递推求解
for(int i=n-2;i>=0;i--){
    for(int j=0;j<=i;j++){
        f[i][j]=Math.min(f[i+1][j],f[i+1][j])+A[i][j]
    }
}
f[0][0]求起点


```

####自顶向下的dp
```
//自顶向下的动态规划
状态定义：
f[i][j]表示从0,0出发,到达i,j的最短路径是什么

//初始化
f[0][0] = A[0][0]

//递推求解
for(int i=1;i<n;i++){
    for(int j=1;j<=i;j++){
        if (i-1.j存在){
            f[i][j]=min(f[i][j],f[i-1][j])
        }
        if (i-1,j-1存在){
            ...
        } //肩上面的点不一定存在
        f[i][j]+=A[i][j]
    }
}
//答案,最后一层最小的。
min(f[n-1][0],f[n-1][1],....)

```

如何想到使用DP

1.满足三个条件：
a.最大/最小问题
b.yes/no
c.Count(*)

2. 无法排序或者交换
3. 如果是所有方案列出来，一定不是动态规划。

### 四要素
1.状态 State
灵感，创造力，存储小规模问题的结果
2. 方程 Function
状态之间的联系，怎么通过小的状态，来算大的状态
3. 初始化 Intialization
最极限的小状态是什么, 起点
4. 答案 Answer
最大的那个状态是什么，终点

### 四种类型
1. Matrix DP (10%)
2. Sequence (40%)
3. Two Sequences DP (40%)
4. Backpack (10%)

### 1. Matrix DP
state: f[x][y] 表示我从起点走到坐标x,y……
function: 研究走到x,y这个点之前的一步
intialize: 起点
answer: 终点
总结如下：
```
state: f[x][y]从起点走到x,y的最短路径
function: f[x][y] = min(f[x-1][y], f[x][y-1]) + A
[x][y]
intialize: f[0][0] = A[0][0]
// f[i][0] = sum(0,0 -> i,0)
// f[0][i] = sum(0,0 -> 0,i)
answer: f[n-1][m-1]
```

```
func minPathSum(grid [][]int) int {
    if grid==nil||len(grid)==0||len(grid[0])==0{
        return 0
    }
   
    row:=len(grid)
    col:=len(grid[0])
    fn:=make([][]int,row) //正确的初始化二维队列的方法
     for i:=0;i<row;i++{
         line:=make([]int,col)
         fn[i]=line
     }
    //init
    /*错误,要累加起来
    fn[0]=grid[0]
    
    for i:=1;i<row;i++{
        fn[i]=make([]int,col)
        fn[i][0]=grid[i][0]
    }
    */
    //正确的方式
    fn[0][0]=grid[0][0]
    for i:=1;i<row;i++{
        fn[i][0]=fn[i-1][0]+grid[i][0]     //累加的方法
    }
    for i:=1;i<col;i++{
        fn[0][i]=fn[0][i-1]+grid[0][i]
    }
    
    
    
    for i:=1;i<row;i++{
        for j:=1;j<col;j++{
            fn[i][j]=min(fn[i][j-1],fn[i-1][j])+grid[i][j]
        }
    }
    
    return fn[row-1][col-1]
}

func min(x,y int) int{
    if x<y{
        return x
    }else{
        return y
    }
}

```

#### unique path
URL:https://leetcode.com/problems/unique-paths/#/description
function: fn(i,j)=fn(i-1,j)+fn(i,j-1)
一把过

init:横竖都是1
```
func uniquePaths(m int, n int) int {
    sum:=make([][]int,m)
    for i:=0;i<m;i++{
        sum[i]=make([]int,n)
        sum[i][0]=1
        
    }
    for i:=0;i<n;i++{
        sum[0][i]=1
    }
    
    for i:=1;i<m;i++{
        for j:=1;j<n;j++{
            sum[i][j]=sum[i-1][j]+sum[i][j-1]
        }
    }
    
    return sum[m-1][n-1]
}

```

#### Unique path 2
URL:

```
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if obstacleGrid==nil{
        return 0
    }
    m:=len(obstacleGrid)
    n:=len(obstacleGrid[0])
    
    sum:=make([][]int,m)
    line:=make([]int,n)
    sum[0]=line
    if obstacleGrid[0][0]==1{  //自己为1，则为不对
        return 0
    }

    //初始化一定做好
    sum[0][0]=1
    
    for i:=1;i<m;i++{
        sum[i]=make([]int,n)

        if obstacleGrid[i][0]!=1{
            sum[i][0]=sum[i-1][0]
        }else{
            sum[i][0]=0
        }
        
        
    }
    for i:=1;i<n;i++{
        if obstacleGrid[0][i]!=1{
            sum[0][i]=sum[0][i-1]
        }else{
            sum[0][i]=0
        }
    }


    
    for i:=1;i<m;i++{
        for j:=1;j<n;j++{
            if obstacleGrid[i][j]!=1{
                sum[i][j]=sum[i-1][j]+sum[i][j-1]
            }else{
                sum[i][j]=0
            }
            
            
        }
    }
    
    return sum[m-1][n-1]
}
```

### 2. Sequence Dp
state: f[i]表示“前i”个位置/数字/字母,(以第i个为)... 

**前5个是怎么样的，前2个是怎么样的，前3个是怎么样的。**

function: f[i] = f[j] … j 是i之前的一个位置
intialize: f[0]..
answer: f[n-1]..

#### 爬梯子
题目：
URL:https://leetcode.com/problems/climbing-stairs/#/description

You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

```
func climbStairs(n int) int {
    if n<=0{
        return 0
    }
    if n==1{
        return 1
    }
    fn:=make([]int,n)
    fn[0]=1
    fn[1]=2
    for i:=2;i<n;i++{
        fn[i]=fn[i-1]+fn[i-2]
    }
    return fn[n-1]
}
```

### jump games

描述：
url:https://leetcode.com/problems/jump-game/#/description

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

For example:
A = [2,3,1,1,4], return true.

A = [3,2,1,0,4], return false.

```
func canJump(nums []int) bool {
    if nums==nil{
        return false
    }
    hash:=make([]int,len(nums))
    
    hash[0]=1   //1.需要关注的点
    
    for i,num:=range nums{
        if hash[i]==0{
            continue
        }
        for j:=1;j<num+1;j++{
          if i+j>=len(nums){
              break //一定要加,越界了
          }
           hash[i+j]=1
        }
    }
    if hash[len(nums)-1]==1{
        return true
    }
    return false
}

```

### lis

Longest Increasing Subsequence
state:
错误的方法: f[i]表示前i个数字中最长的LIS的长度
正确的方法: f[i]表示前i个数字中以第i个结尾的LIS的长
度
function: f[i] = MAX{f[j]+1}, j < i && a[j] <= a
[i])
intialize: f[0..n-1] = 1
answer: max(f[0..n-1])
```
func lengthOfLIS(nums []int) int {
    size:=len(nums)
    record:=make([]int,len(nums))
    for i:=0;i<size;i++{
        record[i]=1
    }
    
    for i:=0;i<size;i++{
        for j:=0;j<i;j++{
            if nums[i]>nums[j]{
                if record[i]<record[j]+1{ //一定要加这个判断,少了这个判断就跪了
                    record[i]=record[j]+1
                }
                
            }
        }
    }
    max:=0
    for i:=0;i<size;i++{
        if max<record[i]{
            max=record[i]
        }
    }
    
    return max
    
}

```

