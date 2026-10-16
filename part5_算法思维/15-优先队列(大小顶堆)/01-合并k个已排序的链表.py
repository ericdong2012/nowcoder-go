"""
https://www.nowcoder.com/practice/65cfde9e5b9b4cf2b6bafa5f3ef33fa6?tpId=196&tqId=37081&ru=/exam/oj

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


def merge_k_lists(lists):
    """write code here"""
    n = len(lists)
    if n == 0:
        return None
    interval = 1
    while interval < n:
        # 如果是2个  0,1 merge即可
        # 如果是3个， 0,1 merge完成后   跳出然后 0，2 merge  所以是interval * 2
        # 如果是4个， 0,1 merge完成后   2，3 merge, 跳出 0,2 merge
        for i in range(0, n - interval, interval * 2):
            lists[i] = merge2_lists(lists[i], lists[i + interval])
        interval *= 2

    return lists[0]


# 2个升序链表的排序
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


def build_linked_list(values):
    """辅助函数：用列表构造链表，返回头节点"""
    dummy = ListNode()
    cur = dummy
    for v in values:
        cur.next = ListNode(v)
        cur = cur.next
    return dummy.next


def print_linked_list(head):
    """辅助函数：打印链表"""
    res = []
    while head is not None:
        res.append(head.val)
        head = head.next
    print(res)


if __name__ == "__main__":
    # [{1,2},{1,4,5},{6}]  ->  {1,1,2,4,5,6}
    lists = [
        build_linked_list([1, 2]),
        build_linked_list([1, 4, 5]),
        build_linked_list([6]),
    ]
    result = merge_k_lists(lists)
    print_linked_list(result)
