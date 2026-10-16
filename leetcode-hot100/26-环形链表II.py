"""142. 环形链表 II（找入环点）
题目：返回环的入口节点；无环返回 None。
思路：Floyd 判圈。相遇后把一个指针放回头部，两个同速前进，
     再次相遇点即入环点。
     证明：头到入环 a，入环到相遇 b，环长 c。相遇时慢走 a+b，
     快走 a+b+kc；快是慢的 2 倍 → a+b = kc → a = kc-b，
     即从头走 a 步与从相遇点走 (k-1)c+b 步同时到入环口。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        slow = fast = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow is fast:                  # 相遇：有环
                p = head
                while p is not slow:
                    p = p.next
                    slow = slow.next
                return p
        return None


if __name__ == "__main__":
    s = Solution()
    a = ListNode(3); b = ListNode(2); c = ListNode(0); d = ListNode(-4)
    a.next, b.next, c.next, d.next = b, c, d, b
    assert s.detectCycle(a) is b
    x = ListNode(1); y = ListNode(2)
    x.next = y
    assert s.detectCycle(x) is None
    print("26 OK")
