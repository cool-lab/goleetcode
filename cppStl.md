### string类
s.size()            字符串长度
s.c_str()       返回char *类型的字符串
s.find(“hi”)        查找子串，返回起始位置，不存在返回string::npos
s.find_last_of(“hi”)    特定方式查找子串
s.substr(0, 5)      获取子串
s.substr(5)

STL容器实现了一些常用的数据结构
1. vector          数组/堆栈
2. list            链表
3. queue           队列
4. deque           双向队列
5. set         集合
6. map         字典、映射
7. priority_queue  优先队列

### Vector
```
#include <vector>
动态数组
vector<int> arr;
arr = vector<int>(10);
arr.push_back(3);
cout << arr[2] << endl;
arr = arr2;
if (arr < arr2)
```

常用的操作
```
arr.clear();
arr.push_back(5);
arr.pop_back();
arr.front();
arr.back();
arr.erase(arr.begin() + 5);
arr.erase(arr.begin(), arr.end());
arr.insert(arr.begin() + 5, 4);

vector的初始化
vector<int> arr(100, 0);
或者
arr = vector<int>(100, 0);
```


### list
```
#include <list>
双向链表
push_back, pop_back, push_front, pop_front
由链表的特性可知，list的迭代器不支持算数运算。例如l.begin() + 5是非法的，而list也不支持[]运算符，因此遍历list只能使用Iterator方式。


```

### deque
```
#include <deque>
双向队列
deque是vector和list的中间产物，其内部结构是若干小段被连接起来的连续空间。
deque支持push_front和pop_front
deque和vector一样支持下标访问
deque的实现比vector和list都要复杂，因此直接影响了它的效率。

```

### priority_queue
```
#include <deque>
用最大堆实现的优先队列
priority_queue中的元素需要支持<操作符。
priority_queue不支持clear操作。
常用操作
pq.top();
pq.push(5);
pq.pop();
```

### set
```
#include <set>
有序的元素集合，set中的元素也要支持<操作符
支持元素插入、删除和查找，复杂度为O(logN)
双向迭代器，不支持随机访问
常用操作
set<int> s;
s.insert(3);
s.erase(5);
if (s.find(5) == s.end())   // 5不在集合中

```

### map
```
map的遍历
for (map<string, int>::iterator mi = m.begin(); mi != m.end(); ++mi) {
cout << mi->first << “ “ << mi->second << endl;
}

```

### 算法
```
#include <algorithm>
sort(arr.begin(), arr.end());
stable_sort(arr.begin(), arr.end());
reverse(arr.begin(), arr.end());
next_permutation(arr.begin(), arr.end());
unique(arr.begin(), arr.end());
```


