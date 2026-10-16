"""
BM12 单链表的排序
https://www.nowcoder.com/practice/f23604257af94d939848729b1a5cda08?tpId=295&tqId=1008897&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个节点数为n的无序单链表，对其按升序排序。

示例1
输入：{1,3,2,4,5}
返回值：{1,2,3,4,5}

示例2
输入：{-1,0,-2}
返回值：{-2,-1,0}
"""


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


# 对比看 链表中的节点每k个一组翻转
# 断点调试，看迭代的过程，很典型的案例
def sort_in_list(head):
    # write code here
    # 可将其拆分成两个链表，然后分别排序，最后转换成合并两个排序链表
    # 怎么找中间元素呢？我们也可以使用快慢双指针，快指针每次两步，慢指针每次一步，那么快指针到达链表尾的时候，慢指针正好走了快指针距离的一半，为中间元素。
    """
    step 1：首先判断链表为空或者只有一个元素，直接就是有序的。
    step 2：准备三个指针，快指针right每次走两步，慢指针mid每次走一步，前序指针left每次跟在mid前一个位置。三个指针遍历链表，当快指针到达链表尾部的时候，慢指针mid刚好走了链表的一半，正好是中间位置。
    step 3：从left位置将链表断开，刚好分成两个子问题开始递归。
    step 4：将子问题得到的链表合并，参考合并两个有序链表。
    """
    return merge_sort(head)


def merge_sort(head):
    if head is None or head.next is None:
        return head
    # 必须要pre， 否则内存超过限制
    # pre是作为中间节点. fast走到最后，pre走到一半
    pre = None
    slow = head
    fast = head
    while fast is not None and fast.next is not None:
        pre = slow
        slow = slow.next
        fast = fast.next.next
    # pre的next断开连接，pre和head都会被断开
    # 将链表分割为2部分，一部分在head中，一部分在slow中
    # head 1, 3
    # slow 2, 5, 4
    pre.next = None
    left = merge_sort(head)    # 后续会将head 拆分中1,3
    right = merge_sort(slow)   # 后续会将head 拆分中2. 5 4  ==> 2,4,5
    # 转换成合并2个排序链表
    return mergexx(left, right)


def mergexx(l, r):
    p = ListNode()
    h = p
    while l is not None and r is not None:
        if l.val < r.val:
            h.next = l
            l = l.next
        else:
            h.next = r
            r = r.next
        h = h.next

    if l is None:
        h.next = r

    if r is None:
        h.next = l

    return p.next


def main12():
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
    result = sort_in_list(node)
    print(result.val)
    print(result.next.val)
    print(result.next.next.val)
    print(result.next.next.next.val)
    print(result.next.next.next.next.val)


if __name__ == "__main__":
    main12()
