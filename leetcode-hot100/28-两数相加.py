"""2. 两数相加（链表）
题目：两个非空链表按逆序存非负整数，返回和的逆序链表。
思路：逐位相加 + 进位，头插尾插皆可；处理最后剩余进位。
复杂度：时间 O(max(m,n))，空间 O(1)（不算结果）。
"""

from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode()
        tail = dummy
        carry = 0
        while l1 or l2 or carry:
            total = carry + (l1.val if l1 else 0) + (l2.val if l2 else 0)
            carry, digit = divmod(total, 10)
            tail.next = ListNode(digit)
            tail = tail.next
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
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
    assert to_list(s.addTwoNumbers(build([2, 4, 3]), build([5, 6, 4]))) == [7, 0, 8]
    assert to_list(s.addTwoNumbers(build([9, 9, 9, 9, 9, 9, 9]), build([9, 9, 9, 9]))) \
        == [8, 9, 9, 9, 0, 0, 0, 1]
    print("28 OK")
