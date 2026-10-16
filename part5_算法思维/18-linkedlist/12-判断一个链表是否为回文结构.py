"""
BM13 判断一个链表是否为回文结构
https://www.nowcoder.com/practice/3fed228444e740c8be66232ce8b87c2f?tpId=295&tqId=1008769&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个链表，请判断该链表是否为回文结构。
回文是指该字符串正序逆序完全一致。

示例1
输入：{1}
返回值：true

示例2
输入：{2,1}
返回值：false

说明：2->1

示例3
输入：{1,2,2,1}
返回值：true

说明：1->2->2->1
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# 直接将链表中的值 用数组存起来，然后前后比较，妙啊
def is_pail(head):
    # write code here
    if head is None:
        return False

    arr = []
    while head is not None:
        arr.append(head.val)
        head = head.next

    i = 0
    j = len(arr) - 1
    while i < j:
        if arr[i] != arr[j]:
            return False
        i += 1
        j -= 1
    return True


# def is_pail(head):
#     # write code here
#     temp = ListNode(next=head)
#     slow = temp
#     fast = temp
#     while fast is not None and fast.next is not None:
#         slow = slow.next
#         fast = fast.next.next
#     mid = slow
#     pre = None
#     tmp = None
#     cur = mid
#     while cur is not None:
#         tmp = cur.next
#         cur.next = pre
#         pre = cur
#         cur = tmp
#     tail = pre
#     while head is not None and tail is not None:
#         if head.val != tail.val:
#             return False
#         head = head.next
#         tail = tail.next
#     return True


def main13():
    # test
    node = ListNode(
        1,
        ListNode(
            3,
            ListNode(
                2,
                ListNode(
                    3,
                    ListNode(1, None),
                ),
            ),
        ),
    )

    # 考虑边界情况
    # result = remove_nth_from_end(node, 2)
    # result = remove_nth_from_end(node, 5)
    result = is_pail(node)
    print(result)


if __name__ == "__main__":
    main13()
