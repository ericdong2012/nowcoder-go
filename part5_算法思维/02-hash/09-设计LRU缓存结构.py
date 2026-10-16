"""
NC93 设计LRU缓存结构
https://www.nowcoder.com/practice/5dfded165916435d9defb053c63f1e84

设计 LRU(最近最少使用)缓存结构，该结构在构造时确定大小，假设大小为 capacity，操作次数是 n，并有如下功能:
1. Solution(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
2. get(key)：如果关键字 key 存在于缓存中，则返回 key 对应的 value 值，否则返回 -1。
3. set(key, value)：将记录(key, value)插入该结构，如果关键字 key 已经存在，则变更其数据值 value，
   如果不存在，则向缓存中插入该组 key-value，
   如果 key-value 的数量超过 capacity，弹出最久未使用的 key-value

提示:
1.某个 key 的 set 或 get 操作一旦发生，则认为这个 key 的记录成了最常使用的，然后都会刷新缓存。
2.当缓存的大小超过 capacity 时，移除最不经常使用的记录。
3.返回的 value 都以字符串形式表达，如果是 set，则会输出"null"来表示(不需要用户返回，系统会自动输出)，方便观察
4.函数 set 和 get 必须以 O(1) 的方式运行
"""

# hash + 双链表
#
# get 操作直接通过 hash 得到值
# set 操作有三种情况：
#   1.当元素在 hash 中：直接修改节点值，并提升到最前面    lru 特性
#   2.当容量充足：创建一个节点，插入到前面
#   3.当容量不足，修改 tail 指针指向的元素值，然后通过索引，修改 hash 的 key，
#     最后把 tail 指针放到双链表头部（删除尾部，添加到头部）

from collections import OrderedDict


class Solution:
    def __init__(self, capacity):
        # 1、用一个 map 来指向双链表中的实体，key -> element
        # 2、用一个双链表来存储对应的 element
        # 3、element 里面存储一个 key 和一个 value，方便从 map 中删除值
        # 这里用 OrderedDict 模拟：队尾为最久未使用，队首为最近使用
        self.cache = OrderedDict()
        self.capacity = capacity

    def get(self, key):
        if key in self.cache:
            # 将该元素移动到队首，lru
            self.cache.move_to_end(key, last=False)
            return self.cache[key]
        else:
            return -1

    def set(self, key, value):
        # 1、如果记录已经存在，那么将其移动前面，刷新值(值可能有变化)
        # 2、如果不存在，判断 LRU 长度是否达到最大值，如果长度没有达到最大值，插入元素到头部
        # 3、如果达到最大值，移除末尾元素，插入元素到头部
        if key in self.cache:
            self.cache.move_to_end(key, last=False)
            self.cache[key] = value
        else:
            if len(self.cache) == self.capacity:
                # 获取最后一个元素，从 list 移除
                # OrderedDict.popitem(last=True) 弹出队尾(最久未使用)
                self.cache.popitem(last=True)
            # 未达到长度，将元素放到队首；cache 记录 key->value
            self.cache[key] = value
            self.cache.move_to_end(key, last=False)
