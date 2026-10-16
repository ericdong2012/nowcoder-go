"""
BM11 链表相加(二)
https://www.nowcoder.com/practice/c56f6c70fb3f4849bc56e33ff2a50b6b?tpId=295&tqId=1008772&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
假设链表中每一个节点的值都在 0-9 之间，那么链表整体就可以代表一个整数，给定两个这种链表，请生成代表两个整数相加值的结果链表。
例如：链表1为 9->3->7， 链表2为 6->3， 最后生成新的结果链表为 1->0->0->0

示例1
输入：[9, 3, 7], [6, 3]
返回值：{1,0,0,0}

示例2
输入：[0], [6, 3]
返回值：{6, 3}
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def reverse_link(head1):
    pre = None
    current = head1
    while current is not None:
        next_node = current.next
        current.next = pre
        pre = current
        current = next_node

    return pre


def add_in_list(head1, head2):
    # write code here
    """
    因为两个链表长度可能不一样，从高位到低位运算无法实现，所以首先想到要反转两个单链表，这题因为没有要求要维持原链表结构不变，直接在原链表上翻转即可
    翻转链表之后，依次对两个链表的对齐节点进行运算即可，因为两个链表的长度可能不一样，这里有个小技巧，较短的链表高位都当作0即可，计算可能产生进位，用一个临时变量保存是否有进位
    """
    t1 = reverse_link(head1)
    t2 = reverse_link(head2)
    # 哑结点
    virtual_node = ListNode()
    # 替哑节点干活的
    p = virtual_node
    # 进位
    carry = 0
    while t1 is not None or t2 is not None:
        # 关键点, v 保存了上次的carry值， 同时会不断累加当前的t1, t2的值
        v = carry
        if t1 is not None:
            v += t1.val
            t1 = t1.next
        if t2 is not None:
            v += t2.val
            t2 = t2.next
        # 进位处理
        if v >= 10:
            v -= 10
            carry = 1
        else:
            carry = 0
        new_node = ListNode(v)
        p.next = new_node
        p = p.next
    # 刚好有进位, 需要在整个链的后面加上一个1
    # 1, 2, 4, 9
    # 2, 3, 6
    # 3, 5, 0, 0, 1
    if carry == 1:
        p.next = ListNode(1)
    # 将链表反转回来
    # virtualNode ---> nil
    # p
    result = reverse_link(virtual_node.next)
    return result


def main11():
    # test
    node = ListNode(
        1,
        ListNode(
            2,
            ListNode(
                3,
                ListNode(
                    6,
                    ListNode(7, None),
                ),
            ),
        ),
    )

    node1 = ListNode(
        4,
        ListNode(
            5,
            ListNode(
                6,
                ListNode(7, None),
            ),
        ),
    )

    # 考虑边界情况
    # result = remove_nth_from_end(node, 2)
    # result = remove_nth_from_end(node, 5)
    result = add_in_list(node, node1)
    print(result.val)
    print(result.next.val)
    print(result.next.next.val)
    print(result.next.next.next.val)
    print(result.next.next.next.next.val)


if __name__ == "__main__":
    main11()
