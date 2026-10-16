"""
BM6 判断链表中是否有环
https://www.nowcoder.com/practice/650474f313294468a4ded3ce0f7898b9

描述
判断给定的链表中是否有环。如果有环则返回 true，否则返回 false。

数据范围：链表长度 0 <= n <= 10000，链表中任意节点的值满足 |val| <= 100000
要求：空间复杂度 O(1)，时间复杂度 O(n)

示例1
输入：{3,2,0,-4},1
返回值：true

示例2
输入：{1},-1
返回值：false
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def has_cycle(head):
    # 说明：题目描述比较奇怪，实际要做的是判断链表中有没有环

    # 单链表，空链表
    if head is None or head.next is None:
        return False
    # 快慢指针
    p1 = head
    p2 = head.next

    while p2 is not None and p2.next is not None:
        p1 = p1.next
        p2 = p2.next.next
        if p1 == p2:
            return True

    return False


def main06():
    # test
    node1 = ListNode(
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

    result = has_cycle(node1)
    print(result)


if __name__ == "__main__":
    main06()
