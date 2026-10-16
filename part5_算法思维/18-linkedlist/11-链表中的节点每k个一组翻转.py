"""
BM3 链表中的节点每k个一组翻转
https://www.nowcoder.com/practice/b49c3dc907814e9bbfa8437c251b028e?tpId=295&tqId=722&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
将给出的链表中的节点每 k 个一组翻转，返回翻转后的链表
如果链表中的节点数不是 k 的倍数，将最后剩下的节点保持原样
你不能更改节点中的值，只能更改节点本身。

示例1
输入：{1,2,3,4,5},2
返回值：{2,1,4,3,5}

示例2
输入：{},1
返回值：{}
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# 对比看 单链表的排序
# 对比看 两两交换链表的节点
# 找到尾节点，交换
def reverse_k_group(head, k):
    # write code here
    dummy = ListNode(next=head)
    pre = dummy

    while head is not None:
        tail = pre
        # 滑动窗口, 大小为k，找尾节点
        for i in range(k):
            # tail往前走
            tail = tail.next
            # 长度不足k
            if tail is None:
                return dummy.next
        # 保存下一组的起始节点
        next_node = tail.next
        # 拿到反转后的头和尾(之前链表的尾和头)
        head1, tail1 = my_reverse(head, tail)
        # 做拼接
        pre.next = head1
        tail1.next = next_node
        # 将pre往前移动, 移动到反转后的尾节点， 将head往前移动到下一组的起始节点(pre/tail1的下一个节点)
        pre = tail1
        head = next_node
    return dummy.next


# 真正难理解的是这里
# 参考题解中的c++代码， 画图理解
def my_reverse(head, tail):
    # pre = ListNode(next=None)
    # 1. 头节点(哑节点)，而不是尾节点
    pre = tail.next
    cur = head
    # pre 走到tail 位置，退出
    while pre != tail:
        # 保存下一位结果
        next_node = cur.next
        # 将cur的下一位断开，指向pre
        cur.next = pre
        # 将pre移动到cur, 将cur移动到next
        pre = cur
        cur = next_node

    return tail, head


def main03():
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

    result = reverse_k_group(node, 3)
    print(result.val)
    print(result.next.val)
    print(result.next.next.val)
    print(result.next.next.next.val)
    print(result.next.next.next.next.val)


if __name__ == "__main__":
    main03()
