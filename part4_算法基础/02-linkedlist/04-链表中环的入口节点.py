"""
https://www.nowcoder.com/practice/253d2c59ec3e4bc68da16833f79a38e4

链表中环的入口节点
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def entry_node_of_loop(p_head):
    # 方式一 快慢指针
    # 看题解，看图，有链表，有环 从相遇节点到入口节点的距离 == 从头节点到入口节点的距离（画图，数学分析得来）
    if p_head is None:
        return p_head
    fast, slow = p_head, p_head
    while fast is not None and fast.next is not None:
        # 1、先用快慢指针判断链表是否有环，如果有环找到快慢指针相遇的节点，注意该节点不一定是环的入口节点
        slow = slow.next
        fast = fast.next.next
        if fast == slow:
            # 2. 此时相遇，通过数学分析从相遇节点到入口节点的距离 == 从头节点到入口节点的距离，
            # 所以让慢指针回到表头，此时快慢执行继续往下走，快慢指针再相遇就是入口处
            slow = p_head
            while fast != slow:
                fast = fast.next
                slow = slow.next
            return fast

    return None
