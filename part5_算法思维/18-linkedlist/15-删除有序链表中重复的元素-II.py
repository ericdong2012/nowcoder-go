"""
BM16 删除有序链表中重复的元素-II
https://www.nowcoder.com/practice/71cef9f8b5564579bf7ed93fbe0b2024?tpId=295&tqId=663&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给出一个升序排序的链表，删除链表中的所有重复出现的元素，只保留原链表中只出现一次的元素。

示例1
输入：{1,2,2}
返回值：{1}

示例2
输入：{}
返回值：{}
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def delete_duplicates(head):
    # write code here
    if head is None:
        return head
    res = ListNode()
    # res 的下一位指向了 head
    res.next = head
    # res 的指针，cur 是干活的，是一个活动 Node
    cur = res
    while cur.next is not None and cur.next.next is not None:
        # 遇到相邻两个节点值相同
        if cur.next.val == cur.next.next.val:
            temp = cur.next.next.val
            # 将所有相同的都跳过，中心扩散
            # 1,2,2 第一次进入删除掉了第一个 2，1 指向了后一个 2，第二次进入 cur.next 是后一个 2，
            # 依然满足条件进入循环，cur.next = nil，删除后一个 2，无法再次进入，跳出循环
            while cur.next is not None and cur.next.val == temp:
                # 跳过 cur.next（cur, cur.next, cur.next.next）
                cur.next = cur.next.next
        else:
            cur = cur.next
    # 返回时去掉表头，初始化时 res 的下一位指向了 head
    return res.next


def main16():
    # test
    node = ListNode(
        1,
        ListNode(
            2,
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
    # dummyult = remove_nth_from_end(node, 2)
    # dummyult = remove_nth_from_end(node, 5)
    dummyult = delete_duplicates(node)
    # 打印整个链表（遍历到 None 为止，原 Go 版本逐层打印节点，nil 安全；Python 这里改为遍历打印）
    cur = dummyult
    while cur is not None:
        print(cur.val)
        cur = cur.next


if __name__ == "__main__":
    main16()
