"""
BM14 链表的奇偶重排
https://www.nowcoder.com/practice/02bf49ea45cd486daa031614f9bd6fc3?tpId=295&tqId=1073463&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个单链表，请设定一个函数，将链表的奇数位节点和偶数位节点分别放在一起，重排后输出。
注意是节点的编号而非节点的数值

示例1
输入：{1,2,3,4,5,6}
返回值：{1,3,5,2,4,6}

说明：
1->2->3->4->5->6->NULL
重排后为
1->3->5->2->4->6->NULL

示例2
输入：{1,4,6,3,7}
返回值：{1,6,7,4,3}
说明：
1->4->6->3->7->NULL
重排后为
1->6->7->4->3->NULL
奇数位节点有1,6,7，偶数位节点有4,3。重排后为1,6,7,4,3
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def odd_even_list(head):
    # write code here
    if head is None or head.next is None:
        return head
    # 奇数  1,3,2,5,4
    odd_head = head
    # 偶数  3,2,5,4
    even_head = head.next
    # 替evenHead干活的， 如果用evenHead 直接执行， 最后evenHead 停留在最后一个偶数位，只有一个节点，所以需要找一个替他干活的
    even = even_head
    # 将下一位保存起来
    next_node = head.next.next
    count = 0
    while next_node is not None:
        count += 1
        if count % 2 == 1:
            odd_head.next = next_node
            odd_head = odd_head.next
            # 1，2，5，4
            # 1, 2, 4
        else:
            # 此处的next已经在下面的 next = next.next 往下移动过了
            even.next = next_node
            even = even.next
            # 3，5，4
        next_node = next_node.next

    # 偶数的尾巴断开， evenHead也随即断开
    even.next = None
    # 奇数指向偶数
    odd_head.next = even_head
    return head


def main14():
    # test
    node = ListNode(
        1,
        ListNode(
            3,
            ListNode(
                2,
                ListNode(
                    5,
                    ListNode(4, None),
                ),
            ),
        ),
    )

    # 考虑边界情况
    # result = remove_nth_from_end(node, 2)
    # result = remove_nth_from_end(node, 5)
    result = odd_even_list(node)
    print(result.val)
    print(result.next.val)
    print(result.next.next.val)
    print(result.next.next.next.val)
    print(result.next.next.next.next.val)


if __name__ == "__main__":
    main14()
