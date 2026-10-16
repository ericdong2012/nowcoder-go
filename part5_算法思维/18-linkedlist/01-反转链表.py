"""
BM1 反转链表
https://www.nowcoder.com/practice/75e878df47f24fdc9dc3e400ec6058ca?tpId=295&tqId=23286&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述：
给定一个单链表的头结点pHead(该头节点是有值的，比如在下图，它的val是1)，长度为n，反转该链表后，返回新链表的表头。

数据范围： 0 <= n <= 1000

要求：空间复杂度 O(1) ，时间复杂度 O(n)

如当输入链表{1,2,3}时，经反转后，原链表变为{3,2,1}，所以对应的输出为{3,2,1}

思路：
1. 双指针 / 手动栈
    1）pre指针指向已经反转好的链表的最后一个节点，最开始没有反转，所以指向nullptr
    2）cur指针指向待反转链表的第一个节点，最开始第一个节点待反转，所以指向head
    3）nex指针指向待反转链表的第二个节点，目的是保存链表，因为cur改变指向后，后面的链表则失效了，所以需要保存

    接下来，循环执行以下三个操作
    1）nex = cur->next, 保存作用
    2）cur->next = pre 未反转链表的第一个节点的下个指针指向已反转链表的最后一个节点
    3）pre = cur， cur = nex; 指针后移，操作下一个未反转链表的第一个节点
    循环条件，当然是cur != nullptr
    循环结束后，cur当然为nullptr，所以返回pre，即为反转后的头结点

2. 递归
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def reverse_list(p_head):
    # 3指针
    # if p_head is None:
    #     return p_head
    # pre = None
    # cur = p_head
    # # 用来保存结果，最终保存的是反转后链表的头节点
    # next_node = None
    # while cur is not None:
    #     next_node = cur.next
    #     cur.next = pre    # 将原来cur.next 指向前一个节点，实现反转
    #     pre, cur = cur, next_node
    # return next_node

    # 最优解，单指针
    if p_head is None or p_head.next is None:
        return p_head

    new_head = None
    while p_head is not None:
        p_next = p_head.next
        p_head.next = new_head
        new_head = p_head
        p_head = p_next
    return new_head

    # 递归
    # if p_head is None or p_head.next is None:
    #     return p_head
    #
    # tmp = reverse_list(p_head.next)
    # p_head.next.next = p_head
    # p_head.next = None
    # return tmp


if __name__ == "__main__":
    # test
    node = ListNode(
        1,
        ListNode(
            2,
            ListNode(
                3,
                ListNode(
                    4,
                    ListNode(5, None),
                ),
            ),
        ),
    )

    result = reverse_list(node)
    print(result.val)
    print(result.next.val)
    print(result.next.next.val)
    print(result.next.next.next.val)
    print(result.next.next.next.next.val)
