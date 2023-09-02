package _3_hash

import "container/list"

/*
NC93 设计LRU缓存结构
https://www.nowcoder.com/practice/5dfded165916435d9defb053c63f1e84?tpId=117&tqId=37804&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=

设计LRU(最近最少使用)缓存结构，该结构在构造时确定大小，假设大小为 capacity ，操作次数是 n ，并有如下功能:
1. Solution(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
2. get(key)：如果关键字 key 存在于缓存中，则返回key对应的value值，否则返回 -1 。
3. set(key, value)：将记录(key, value)插入该结构，如果关键字 key 已经存在，则变更其数据值 value，
    如果不存在，则向缓存中插入该组 key-value ，
    如果key-value的数量超过capacity，弹出最久未使用的key-value

提示:
1.某个key的set或get操作一旦发生，则认为这个key的记录成了最常使用的，然后都会刷新缓存。
2.当缓存的大小超过capacity时，移除最不经常使用的记录。
3.返回的value都以字符串形式表达，如果是set，则会输出"null"来表示(不需要用户返回，系统会自动输出)，方便观察
4.函数set和get必须以O(1)的方式运行
5.为了方便区分缓存里key与value，下面说明的缓存里key用""号包裹

示例1
输入：
["set","set","get","set","get","set","get","get","get"],[[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]],2
返回值：
["null","null","1","null","-1","null","-1","3","4"]

说明：
我们将缓存看成一个队列，最后一个参数为2代表capacity，所以
Solution s = new Solution(2);
s.set(1,1); //将(1,1)插入缓存，缓存是{"1"=1}，set操作返回"null"
s.set(2,2); //将(2,2)插入缓存，缓存是{"2"=2，"1"=1}，set操作返回"null"
output=s.get(1);// 因为get(1)操作，缓存更新，缓存是{"1"=1，"2"=2}，get操作返回"1"
s.set(3,3); //将(3,3)插入缓存，缓存容量是2，故去掉某尾的key-value，缓存是{"3"=3，"1"=1}，set操作返回"null"
output=s.get(2);// 因为get(2)操作，不存在对应的key，故get操作返回"-1"
s.set(4,4); //将(4,4)插入缓存，缓存容量是2，故去掉某尾的key-value，缓存是{"4"=4，"3"=3}，set操作返回"null"
output=s.get(1);// 因为get(1)操作，不存在对应的key，故get操作返回"-1"
output=s.get(3);//因为get(3)操作，缓存更新，缓存是{"3"=3，"4"=4}，get操作返回"3"
output=s.get(4);//因为get(4)操作，缓存更新，缓存是{"4"=4，"3"=3}，get操作返回"4"

*/

// https://blog.csdn.net/chenbaoke/article/details/42780895
/*
hash + 双链表

get操作直接通过hash得到值
set操作有三种情况：
	1.当元素在hash中： 直接修改节点值，并提升到最前面      lru 特性
	2.当容量充足： 创建一个节点，插入到前面
	3.当容量不足， 修改tail指针指向的元素值，然后通过索引，修改hash的key，最后把tail指针放到双链表头部（删除尾部，添加到头部）

*/

type element struct {
	key   int
	value int
}

type Solution struct {
	// 1、用一个map来指向双链表中的实体，key， element
	// 2、用一个双链表来存储对应的element
	// 3、element里面存储一个key和一个value， 方便从map中删除值
	cache    map[int]*list.Element
	ll       *list.List
	capacity int
	size     int
}

func New(capacity int) Solution {
	return Solution{
		cache:    make(map[int]*list.Element),
		ll:       list.New(),
		capacity: capacity,
		size:     0,
	}
}

func (this *Solution) get(key int) int {
	if e, ok := this.cache[key]; ok {
		// 将尾部元素移动到头部, lru
		this.ll.MoveToFront(e)
		kv := e.Value.(*element)
		return kv.value
	} else {
		return -1
	}
}

func (this *Solution) set(key int, value int) {
	// 1、如果记录已经存在，那么将其移动前面，刷新值(值可能有变化)
	// 2、如果不存在， 判断LRU长度是否达到最大值, 如果长度没有达到最大值, 插入元素到头部
	// 3、如果达到最大值，移除末尾元素, 插入元素到头部
	if e, ok := this.cache[key]; ok {
		this.ll.MoveToFront(e)
		kv := e.Value.(*element)
		kv.value = value
	} else {
		if this.size == this.capacity {
			// 获取最后一个元素， 从list移除
			tail := this.ll.Back()
			this.ll.Remove(tail)
			// 从cache中移除
			kv := tail.Value.(*element)
			delete(this.cache, kv.key)
			// 长度减一
			this.size--
		}
		// 未达到长度， 将元素放到list头部; cache记录 key, e ; 长度+1
		e := this.ll.PushFront(&element{key: key, value: value})
		this.cache[key] = e
		this.size++
	}
}
