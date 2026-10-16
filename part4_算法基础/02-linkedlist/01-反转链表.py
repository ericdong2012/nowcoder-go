"""
BM1 反转链表
https://www.nowcoder.com/practice/75e878df47f24fdc9dc3e400ec6058ca

描述
给定一个单链表的头结点 pHead(该头节点是有值的，比如在下图，它的 val 是 1)，长度为 n，
反转该链表后，返回新链表的表头。

数据范围：0 <= n <= 1000
要求：空间复杂度 O(1)，时间复杂度 O(n)

如当输入链表 {1,2,3} 时，经反转后，原链表变为 {3,2,1}，所以对应的输出为 {3,2,1}

思路：
1. 双指针 / 手动栈
    1）pre 指针指向已经反转好的链表的最后一个节点，最开始没有反转，所以指向 nullptr
    2）cur 指针指向待反转链表的第一个节点，最开始第一个节点待反转，所以指向 head
    3）next 指针指向待反转链表的第二个节点，目的是保存链表，因为 cur 改变指向后，后面的链表则失效了，所以需要保存

    接下来，循环执行以下三个操作
    1）next = cur.next, 保存作用
    2）cur.next = pre 未反转链表的第一个节点的下个指针指向已反转链表的最后一个节点
    3）pre = cur，cur = next；指针后移，操作下一个未反转链表的第一个节点
    循环条件，当然是 cur != nullptr
    循环结束后，cur 当然为 nullptr，所以返回 pre，即为反转后的头结点

2. 递归

tips:
    看题 5 分钟，思考不出来，直接看题解(讨论 + 排行，根据语言筛选)，优先看图解，然后再看代码，最后自己实现。
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def reverse_list(p_head):
    """
    最优解，单指针
    """
    # 3 指针
    # if p_head is None:
    #     return p_head
    # pre = None
    # cur = p_head
    # # 用来保存结果，最终保存的是反转后链表的头节点
    # nxt = None
    # while cur is not None:
    #     nxt = cur.next
    #     cur.next = pre  # 将原来 cur.next 指向前一个节点，实现反转
    #     pre, cur = cur, nxt
    # return nxt

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
    # tmp = reverse_list(p_head.next)
    # p_head.next.next = p_head
    # p_head.next = None
    # return tmp


def main():
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


if __name__ == "__main__":
    main()
