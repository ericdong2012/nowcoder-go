"""23. 合并 K 个升序链表
题目：合并 k 条升序链表。
思路 A：小顶堆存每条链的当前节点，每次弹最小接到结果，弹出的 next 入堆。
思路 B：分治两两合并（同 148 的 merge）。
复杂度：堆版时间 O(N log k)（N 总节点数），空间 O(k)。
"""

from typing import List, Optional
import heapq


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    # 让 heapq 可以比较（否则同值节点比较会报错）
    def __lt__(self, other):
        return self.val < other.val


class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        dummy = ListNode()
        tail = dummy
        heap = [l for l in lists if l]
        heapq.heapify(heap)
        while heap:
            node = heapq.heappop(heap)
            tail.next = node
            tail = tail.next
            if node.next:
                heapq.heappush(heap, node.next)
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
    out = s.mergeKLists([build([1, 4, 5]), build([1, 3, 4]), build([2, 6])])
    assert to_list(out) == [1, 1, 2, 3, 4, 4, 5, 6]
    assert s.mergeKLists([]) is None
    print("34 OK")
