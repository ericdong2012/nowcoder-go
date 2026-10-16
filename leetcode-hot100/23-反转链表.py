"""206. 反转链表（链表基本功）
题目：反转单链表。
思路：迭代三指针 prev/cur/next 逐个掉头；递归版一行思想：
     head.next.next = head。
复杂度：时间 O(n)，空间 O(1)（迭代）。
"""

from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev, cur = None, head
        while cur:
            nxt = cur.next
            cur.next = prev
            prev = cur
            cur = nxt
        return prev


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
    assert to_list(s.reverseList(build([1, 2, 3, 4, 5]))) == [5, 4, 3, 2, 1]
    assert to_list(s.reverseList(build([]))) == []
    print("23 OK")
