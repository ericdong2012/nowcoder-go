"""146. LRU 缓存（哈希 + 双向链表）
题目：O(1) 的 get/put，容量满时淘汰最久未使用。
思路：哈希表 {key: node} + 双向链表维护使用顺序。
     头部是最新；get/put 都把节点移到头部；满了删尾部。
     Python 可用 OrderedDict 一行实现，面试建议手写双向链表版。
复杂度：两个操作均 O(1)。
"""


class Node:
    __slots__ = ("key", "val", "prev", "next")

    def __init__(self, key=0, val=0):
        self.key, self.val = key, val
        self.prev = self.next = None


class LRUCache:
    def __init__(self, capacity: int):
        self.cap = capacity
        self.map = {}                                   # key -> Node
        self.head, self.tail = Node(), Node()           # 哑头尾
        self.head.next = self.tail
        self.tail.prev = self.head

    def _remove(self, node: Node) -> None:
        node.prev.next = node.next
        node.next.prev = node.prev

    def _add_front(self, node: Node) -> None:
        node.next = self.head.next
        node.prev = self.head
        self.head.next.prev = node
        self.head.next = node

    def get(self, key: int) -> int:
        if key not in self.map:
            return -1
        node = self.map[key]
        self._remove(node)
        self._add_front(node)
        return node.val

    def put(self, key: int, value: int) -> None:
        if key in self.map:
            node = self.map[key]
            node.val = value
            self._remove(node)
            self._add_front(node)
            return
        if len(self.map) >= self.cap:
            lru = self.tail.prev
            self._remove(lru)
            del self.map[lru.key]
        node = Node(key, value)
        self.map[key] = node
        self._add_front(node)


if __name__ == "__main__":
    c = LRUCache(2)
    c.put(1, 1)
    c.put(2, 2)
    assert c.get(1) == 1
    c.put(3, 3)              # 淘汰 2
    assert c.get(2) == -1
    c.put(4, 4)              # 淘汰 1
    assert c.get(1) == -1
    assert c.get(3) == 3
    assert c.get(4) == 4
    print("35 OK")
