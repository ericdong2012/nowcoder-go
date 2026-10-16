"""141. 环形链表（Floyd 判圈）
题目：判断链表是否有环。
思路：快慢指针，快走 2 慢走 1；有环必相遇，无环 fast 先到 None。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        slow = fast = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow is fast:
                return True
        return False


if __name__ == "__main__":
    s = Solution()
    a = ListNode(3); b = ListNode(2); c = ListNode(0); d = ListNode(-4)
    a.next, b.next, c.next, d.next = b, c, d, b
    assert s.hasCycle(a) is True
    x = ListNode(1); y = ListNode(2)
    x.next = y
    assert s.hasCycle(x) is False
    print("25 OK")
