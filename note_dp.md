
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



