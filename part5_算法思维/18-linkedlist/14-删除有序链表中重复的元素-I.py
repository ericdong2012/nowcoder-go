"""
BM15 删除有序链表中重复的元素-I
https://www.nowcoder.com/practice/c087914fae584da886a0091e877f2c79?tpId=295&tqId=664&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
删除给出链表中的重复元素（链表中元素从小到大有序），使链表中的所有元素都只出现一次
例如：
给出的链表为 1->1->2，返回 1->2.
给出的链表为 1->1->2->3->3，返回 1->2->3.

示例1
输入：{1,1,2}
返回值：{1,2}

示例2
输入：{}
返回值：{}
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# 思路： 有序链表，所以如果是有重复的，必然连着的
def delete_duplicates1(head):
    # 链表是有序的
    if head is None:
        return head
    # pre 是指向 head 的，干活的，最终返回的还是 head
    pre = head
    while pre.next is not None:
        if pre.val == pre.next.val:
            pre.next = pre.next.next
        else:
            pre = pre.next

    return head


def main15():
    # test
    node = ListNode(
        1,
        ListNode(
            1,
            ListNode(
                2,
                ListNode(
                    4,
                    ListNode(5, None),
                ),
            ),
        ),
    )

    # 考虑边界情况
    # result = remove_nth_from_end(node, 2)
    # result = remove_nth_from_end(node, 5)
    result = delete_duplicates1(node)
    # 打印整个链表（遍历到 None 为止，原 Go 版本逐层打印节点，nil 安全；Python 这里改为遍历打印）
    cur = result
    while cur is not None:
        print(cur.val)
        cur = cur.next


if __name__ == "__main__":
    main15()
