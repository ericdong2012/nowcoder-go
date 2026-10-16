"""
https://www.nowcoder.com/practice/65cfde9e5b9b4cf2b6bafa5f3ef33fa6

合并 k 个升序的链表并将结果作为一个升序的链表返回其头节点。

输入：
[{1,2,3},{4,5,6,7}]
返回值：
{1,2,3,4,5,6,7}


输入：
[{1,2},{1,4,5},{6}]
返回值：
{1,1,2,4,5,6}
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# part2_算法思维/18-linkedlist/05-合并k个已排序的链表.go
# 和 合并两个有序数组， 合并区间对比、
# 主要思想就是 迭代
def merge_k_lists(lists):
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


def merge2_lists(l1, l2):
    # 边界处理
    if l1 is None:
        return l2
    if l2 is None:
        return l1
    # 比较两个链表的头结点，将小的作为新列表的头结点
    if l1.val < l2.val:
        head = l1
        head.next = merge2_lists(l1.next, l2)
    else:
        head = l2
        head.next = merge2_lists(l1, l2.next)

    return head


if __name__ == "__main__":
    # [{1,2},{1,4,5},{6}]
    l1 = ListNode(1, ListNode(2))
    l2 = ListNode(1, ListNode(4, ListNode(5)))
    l3 = ListNode(6)
    head = merge_k_lists([l1, l2, l3])
    result = []
    while head is not None:
        result.append(head.val)
        head = head.next
    print(result)
