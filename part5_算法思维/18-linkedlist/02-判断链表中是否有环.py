"""
BM6 判断链表中是否有环
https://www.nowcoder.com/practice/650474f313294468a4ded3ce0f7898b9?tpId=295&tqId=605&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
判断给定的链表中是否有环。如果有环则返回true，否则返回false。

数据范围：链表长度 0 <= n <= 10000，链表中任意节点的值满足 |val| <= 100000
要求：空间复杂度 O(1)，时间复杂度 O(n)

输入分为两部分，第一部分为链表，第二部分代表是否有环，然后将组成的head头结点传入到函数里面。-1代表无环，其它的数字代表有环，这些参数解释仅仅是为了方便读者自测调试。实际在编程时读入的是链表的头节点。

例如输入{3,2,0,-4},1时，对应的链表结构如下图所示：
可以看出环的入口结点为从头结点开始的第1个结点（注：头结点为第0个结点），所以输出true。

示例1
输入：{3,2,0,-4},1
返回值：true

示例2
输入：{1},-1
返回值：false

示例3
输入：{-1,-7,7,-4,19,6,-9,-5,-2,-5},6
返回值：true
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def has_cycle(head):
    # 说明： 题目描述比较奇怪，实际要做的是判断链表中有没有环

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
    # test，画图
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
