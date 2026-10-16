"""138. 随机链表的复制（原地插值法）
题目：深拷贝带 random 指针的链表。
思路（O(1) 额外空间）：把每个拷贝节点插到原节点后面 →
      拷贝.random = 原.random.next（拷贝就在原的旁边）→ 拆成两条链。
复杂度：时间 O(n)，空间 O(1)（不算结果）。
"""


class Node:
    def __init__(self, x: int, next: "Node" = None, random: "Node" = None):
        self.val = x
        self.next = next
        self.random = random


class Solution:
    def copyRandomList(self, head: "Node") -> "Node":
        if not head:
            return None
        # 1) 每个原节点后面插入拷贝
        cur = head
        while cur:
            cur.next = Node(cur.val, cur.next)
            cur = cur.next.next
        # 2) 设置拷贝节点的 random
        cur = head
        while cur:
            if cur.random:
                cur.next.random = cur.random.next   # 拷贝的random = 原random的拷贝
            cur = cur.next.next
        # 3) 拆链：恢复原链 + 收出拷贝链
        dummy = Node(0)
        tail, cur = dummy, head
        while cur:
            copy = cur.next
            cur.next = copy.next
            tail.next = copy
            tail = tail.next
            cur = cur.next
        return dummy.next


def build(spec):
    """spec: [(val, random_index_or_None), ...]，返回 (head, nodes)。"""
    nodes = [Node(v) for v, _ in spec]
    for i, (_, rnd) in enumerate(spec):
        if i + 1 < len(nodes):
            nodes[i].next = nodes[i + 1]
        nodes[i].random = nodes[rnd] if rnd is not None else None
    return nodes[0] if nodes else None, nodes


if __name__ == "__main__":
    s = Solution()
    spec = [(7, None), (13, 0), (11, 4), (10, 2), (1, 0)]
    head, nodes = build(spec)
    copy = s.copyRandomList(head)
    # 收集拷贝链，逐位验证值与 random 指向（random 应指向"拷贝链中对应下标"）
    copies = []
    c = copy
    while c:
        copies.append(c)
        c = c.next
    assert len(copies) == len(spec)
    for i, (v, rnd) in enumerate(spec):
        assert copies[i] is not nodes[i]        # 深拷贝，不共享节点
        assert copies[i].val == v
        expected = copies[rnd] if rnd is not None else None
        assert copies[i].random is expected
    # 原链应保持原状（next 已恢复）
    a = head
    for v, _ in spec:
        assert a.val == v
        a = a.next
    print("32 OK")
