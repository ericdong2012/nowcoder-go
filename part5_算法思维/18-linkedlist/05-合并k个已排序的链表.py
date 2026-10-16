"""
BM5 合并k个已排序的链表
https://www.nowcoder.com/practice/65cfde9e5b9b4cf2b6bafa5f3ef33fa6?tpId=295&tqId=724&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
合并 k 个升序的链表并将结果作为一个升序的链表返回其头节点。

示例1
输入：[{1,2,3},{4,5,6,7}]
返回值：{1,2,3,4,5,6,7}

示例2
输入：[{1,2},{1,4,5},{6}]
返回值：{1,1,2,4,5,6}
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# part2_算法思维/12-分治/01-合并K个已排序的链表.go
def merge_k_lists(lists):
    # write code here
    """
    方法一： 在合并两个排序链表的基础上，加上分治， 变成合并k个已排序的链表
    方法二： 优先队列
    """
    if len(lists) == 0:
        return None

    if len(lists) == 1:
        return lists[0]
    # 将最前面的两个先合并
    cur = merge2_lists(lists[0], lists[1])
    # 将后面的和前面的合并
    for i in range(2, len(lists)):
        cur = merge2_lists(cur, lists[i])

    return cur


# 合并2个链表  解法1
# def merge2_lists(l1, l2):
#     # 边界处理
#     if l1 is None:
#         return l2
#     if l2 is None:
#         return l1
#     # 比较两个链表的头结点，将小的作为新列表的头结点
#     head = None
#     if l1.val < l2.val:
#         head = l1
#         head.next = merge2_lists(l1.next, l2)
#     else:
#         head = l2
#         head.next = merge2_lists(l1, l2.next)
#
#     return head


# 合并2个链表  解法2
def merge2_lists(l1, l2):
    dummy_head = ListNode()
    head = dummy_head
    while l1 is not None and l2 is not None:
        if l1.val < l2.val:
            head.next = l1
            l1 = l1.next
        else:
            head.next = l2
            l2 = l2.next
        head = head.next
    # 对剩余的做处理（l1, l2中至少有一个为空，剩余的是大数）
    if l1 is not None:
        head.next = l1
    if l2 is not None:
        head.next = l2
    return dummy_head.next


# python  优先队列解法
# from queue import PriorityQueue
# def mergeKLists(lists):
#     # 小顶堆
#     pq = PriorityQueue()
#     # 遍历所有链表第一个元素
#     for i in range(len(lists)):
#         # 不为空则加入小顶堆
#         if lists[i] is not None:
#             pq.put((lists[i].val, i))
#             lists[i] = lists[i].next
#     # 加一个表头
#     res = ListNode(-1)
#     head = res
#     # 直到小顶堆为空
#     while not pq.empty():
#         # 取出最小的元素
#         val, idx = pq.get()
#         # 连接
#         head.next = ListNode(val)
#         head = head.next
#         if lists[idx] is not None:
#             # 每次取出链表的后一个元素加入小顶堆
#             pq.put((lists[idx].val, idx))
#             lists[idx] = lists[idx].next
#     return res.next


def main05():
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

    node2 = ListNode(
        2,
        ListNode(
            3,
            ListNode(
                5,
                ListNode(
                    7,
                    ListNode(9, None),
                ),
            ),
        ),
    )

    result = merge_k_lists([node1, node2])
    print(result.val)
    print(result.next.val)
    print(result.next.next.val)
    print(result.next.next.next.val)
    print(result.next.next.next.next.val)
    print(result.next.next.next.next.next.val)
    print(result.next.next.next.next.next.next.val)
    print(result.next.next.next.next.next.next.next.val)
    print(result.next.next.next.next.next.next.next.next.val)
    print(result.next.next.next.next.next.next.next.next.next.val)


if __name__ == "__main__":
    main05()
