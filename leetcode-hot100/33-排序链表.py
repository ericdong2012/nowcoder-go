"""148. 排序链表（归并排序）
题目：O(n log n) 排序链表。
思路：快慢指针找中点断开 → 递归排序两半 → 合并两个有序链表。
     进阶自底向上迭代可省递归栈。
复杂度：时间 O(n log n)，空间 O(log n)（递归栈）。
"""

from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head or not head.next:
            return head
        # 快慢指针找中点，断开
        slow, fast = head, head.next
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        mid = slow.next
        slow.next = None
        # 递归排序两半后合并
        left = self.sortList(head)
        right = self.sortList(mid)
        return self._merge(left, right)

    def _merge(self, a, b):
        dummy = ListNode()
        tail = dummy
        while a and b:
            if a.val <= b.val:
                tail.next, a = a, a.next
            else:
                tail.next, b = b, b.next
            tail = tail.next
        tail.next = a or b
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
    assert to_list(s.sortList(build([4, 2, 1, 3]))) == [1, 2, 3, 4]
    assert to_list(s.sortList(build([-1, 5, 3, 4, 0]))) == [-1, 0, 3, 4, 5]
    assert to_list(s.sortList(build([]))) == []
    print("33 OK")
