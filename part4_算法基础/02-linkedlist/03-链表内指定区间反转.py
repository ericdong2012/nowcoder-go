"""
https://www.nowcoder.com/practice/b58434e200a648c589ca2063f1faf58c

链表内指定区间反转

输入：{1,2,3,4,5},2,4
返回值：{1,4,3,2,5}

输入：{5},1,1
返回值：{5}

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

@param head ListNode类
@param m int整型
@param n int整型
@return ListNode类
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def reverse_between(head, m, n):
    """write code here"""
    if head is None or head.next is None:
        return head
    # 哑节点
    dummy = ListNode(0, head)
    # 开始和结束前的节点
    pre, end = dummy, dummy
    # 以 {1,2,3,4,5},2,4 为例
    for _ in range(m - 1):
        pre = pre.next

    # pre 在 1，start 到了 2
    start = pre.next
    for _ in range(n):
        end = end.next
    # 结束节点的下一个节点，end 到了 4，next 到了 5
    nxt = end.next
    end.next = None

    # 将开始前的节点指向反转后链表的头节点
    pre.next = reverse(start)
    # 反转后链表的下一个节点是尾节点 next
    start.next = nxt

    # 返回哑节点的下一个节点，即整个链表的开头
    return dummy.next


def reverse(head):
    pre = None
    while head is not None:
        nxt = head.next
        head.next = pre
        pre = head
        head = nxt
    return pre
