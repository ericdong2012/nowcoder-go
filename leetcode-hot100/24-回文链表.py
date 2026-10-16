"""234. 回文链表（反转后半段）
题目：判断链表是否回文，O(n) 时间 O(1) 空间。
思路：快慢指针找中点 → 反转后半段 → 逐个比较 → （可再反转还原）。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        # 快慢指针找中点（slow 停在前半段末尾）
        slow = fast = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        # 反转后半段
        prev = None
        while slow:
            nxt = slow.next
            slow.next = prev
            prev = slow
            slow = nxt
        # 比较
        while prev:
            if prev.val != head.val:
                return False
            prev = prev.next
            head = head.next
        return True


def build(vals):
    dummy = ListNode()
    cur = dummy
    for v in vals:
        cur.next = ListNode(v)
        cur = cur.next
    return dummy.next


if __name__ == "__main__":
    s = Solution()
    assert s.isPalindrome(build([1, 2, 2, 1])) is True
    assert s.isPalindrome(build([1, 2])) is False
    assert s.isPalindrome(build([1])) is True
    print("24 OK")
