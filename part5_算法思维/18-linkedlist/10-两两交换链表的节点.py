"""
NC186 两两交换链表的节点
https://www.nowcoder.com/practice/71f95c23810349f782a1aa6c9bd714b4?tpId=117&tqId=39353&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D196&difficulty=undefined&judgeStatus=undefined&tags=580&title=

描述
给你一个链表，请你两两交换相邻节点，你需要真正交换节点本身，而不是修改节点的值。

两两交换示例：
链表    ：1->2->3->4
交换后  ：2->1->4->3

链表    ：1->2->3
交换后   ：2->1->3

示例1
输入：{1,2,3,4}
返回值：{2,1,4,3}

示例2
输入：{1,2,3}
返回值：{2,1,3}
"""


class ListNode1:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def swap_linked_pair(head):
    if head is None or head.next is None:
        return head
    new_head = head.next
    # 当前节点 -> 下下个节点反转后的结果, 1 -> 4
    head.next = swap_linked_pair(new_head.next)
    # 下一个节点 -> 当前节点   2 -> 1
    new_head.next = head
    return new_head


# 需要调试
# 解法2
# def swap_linked_pair(head):
#     # write code here
#     if head is None or head.next is None:
#         return head
#     ret = head.next
#     while head is not None and head.next is not None:
#         # 将head.next.next保存下来  1,2,3,4 中的3开始的节点保存下来
#         t = head.next.next
#         # head的下一位 反指 向了head（ret指向了head）
#         # 1,2,3,4 中的2指向了1 , ret变成了2,1,4    后面还有一串恶心的东西
#         # 3,4 中的4反指向了3， ret 中变成了2,1,4,3  后面还有一串恶心的东西
#         head.next.next = head
#         # 3,4都存在
#         if t is not None and t.next is not None:
#             # 1指向了4, 原有的head 指向了head.next.next.next
#             head.next = t.next
#         elif t is not None:
#             # t没有下一位，直接指向t, 比如1,2,3  1直接指向3
#             head.next = t
#         else:
#             # t为空，head 断开（3后面的断开）
#             head.next = None
#         # head前移2位
#         # 未移动前是1,4 , 移动后head跑到了3, 现在是3,4
#         head = t
#
#     return ret


if __name__ == "__main__":
    # test
    node = ListNode1(
        1,
        ListNode1(
            2,
            ListNode1(
                3,
                ListNode1(4, None),
            ),
        ),
    )

    result = swap_linked_pair(node)
    print(result.val)
    print(result.next.val)
    print(result.next.next.val)
    print(result.next.next.next.val)
