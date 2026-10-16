"""24. 两两交换链表中的节点
题目：相邻两个节点两两交换，返回新头。
思路：哑结点；每轮让 prev.next 与 prev.next.next 交换：
     prev → b → a → 后续，prev 前进到 a。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(next=head)
        prev = dummy
        while prev.next and prev.next.next:
            a, b = prev.next, prev.next.next
            a.next = b.next
            b.next = a
            prev.next = b
            prev = a
        return dummy.next


def build(vals):
    dummy = ListNode()
    cur = dummy
    for v in vals:
        cur.next = ListNode(v)
        cur = cur.next
    return dummy.next


def to_list(head):
    out = []
    while head:
        out.append(head.val)
        head = head.next
    return out


if __name__ == "__main__":
    s = Solution()
    assert to_list(s.swapPairs(build([1, 2, 3, 4]))) == [2, 1, 4, 3]
    assert to_list(s.swapPairs(build([1]))) == [1]
    assert to_list(s.swapPairs(build([]))) == []
    print("30 OK")
