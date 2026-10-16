"""25. K 个一组翻转链表
题目：每 k 个节点一组翻转，不足 k 个保持原样。
思路：哑结点；对每一组：先确认够 k 个 → 组内翻转 → 接回主链。
      组内翻转即 23 题的局部应用。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        dummy = ListNode(next=head)
        group_prev = dummy
        while True:
            # 检查剩余是否够 k 个
            kth = group_prev
            for _ in range(k):
                kth = kth.next
                if not kth:
                    return dummy.next
            group_next = kth.next
            # 组内翻转：prev 从 group_next 开始（尾接原组后）
            prev, cur = group_next, group_prev.next
            while cur is not group_next:
                nxt = cur.next
                cur.next = prev
                prev = cur
                cur = nxt
            # 接回：group_prev 原来指向组头，现在组头变组尾
            tmp = group_prev.next          # 原组头（新组尾）
            group_prev.next = kth          # 新组头
            group_prev = tmp               # 下一组的 prev = 本组尾


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
    assert to_list(s.reverseKGroup(build([1, 2, 3, 4, 5]), 2)) == [2, 1, 4, 3, 5]
    assert to_list(s.reverseKGroup(build([1, 2, 3, 4, 5]), 3)) == [3, 2, 1, 4, 5]
    print("31 OK")
