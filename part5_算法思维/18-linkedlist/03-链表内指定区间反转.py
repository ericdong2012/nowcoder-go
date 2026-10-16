"""
BM2 链表内指定区间反转
https://www.nowcoder.com/practice/b58434e200a648c589ca2063f1faf58c?tpId=295&tqId=654&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

将一个节点数为 size 链表 m 位置到 n 位置之间的区间反转

输入：{1,2,3,4,5},2,4
返回值： 第4个和第2个发生了反转 3,5没动
{1,4,3,2,5}

输入：{5},1,1
返回值：{5}
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def reverse(head):
    pre = None
    while head is not None:
        next_node = head.next
        head.next = pre
        pre = head
        head = next_node
    return pre


def reverse_between(head, m, n):
    # write code here
    if head is None or head.next is None:
        return head
    # 哑节点
    dummy = ListNode(next=head)
    # 开始和结束前的节点
    pre = dummy
    post = dummy
    for i in range(m - 1):
        pre = pre.next
    for i in range(n):
        post = post.next
    start = pre.next
    # 此时post已经在4， 保存结束节点的下一个节点， 并将结束节点断开
    next_node = post.next
    post.next = None
    # 将开始前的节点指向反转后链表的头节点(post)
    pre.next = reverse(start)
    # 反转后链表的尾节点指向之前断开的节点， 连接上
    start.next = next_node
    # 返回哑节点的下一个节点，即整个链表的开头
    return dummy.next
