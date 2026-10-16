"""19. 删除链表的倒数第 N 个结点（快慢指针）
题目：删除倒数第 n 个节点。
思路：哑结点 + 快指针先走 n 步，快慢同走到快到尾，慢的 next 即目标。
     哑结点统一处理"删头"。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(next=head)
        fast = slow = dummy
        for _ in range(n + 1):          # 快先走 n+1 步，slow 停在目标前一个
            fast = fast.next
        while fast:
            fast = fast.next
            slow = slow.next
        slow.next = slow.next.next
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
    assert to_list(s.removeNthFromEnd(build([1, 2, 3, 4, 5]), 2)) == [1, 2, 3, 5]
    assert to_list(s.removeNthFromEnd(build([1]), 1)) == []
    assert to_list(s.removeNthFromEnd(build([1, 2]), 1)) == [1]
    print("29 OK")
