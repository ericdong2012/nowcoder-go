"""160. 相交链表（双指针）
题目：找两个单链表的交点起始节点（可能不相交）。
思路：双指针各走一条链，走到头换走另一条。若相交，两边走过的总长
     (a+c)+(b+c) 相同，必在交点相遇；不相交则同时到达 None。
复杂度：时间 O(m+n)，空间 O(1)。
"""

from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        a, b = headA, headB
        while a is not b:
            a = a.next if a else headB
            b = b.next if b else headA
        return a


def _build(vals, pos, link=None):
    """pos=-1 不成环；link>0 表示尾部接到该下标节点（造相交）。"""
    nodes = [ListNode(v) for v in vals]
    for i in range(len(nodes) - 1):
        nodes[i].next = nodes[i + 1]
    if pos >= 0:
        nodes[-1].next = nodes[pos]
    return nodes[0] if nodes else None


if __name__ == "__main__":
    s = Solution()
    # 构造相交：A: 1→2→(3→4→5), B: 9→(3→4→5)
    common = _build([3, 4, 5], -1)
    a = ListNode(1); a.next = ListNode(2); a.next.next = common
    b = ListNode(9); b.next = common
    assert s.getIntersectionNode(a, b) is common
    x = _build([1, 2], -1)
    y = _build([3, 4], -1)
    assert s.getIntersectionNode(x, y) is None
    print("22 OK")
