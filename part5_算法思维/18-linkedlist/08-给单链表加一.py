"""
NC189 给单链表加一
https://www.nowcoder.com/practice/a2f1105d2ac5466e9ba8fd61310ba6d1?tpId=117&tqId=39360&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D196&difficulty=undefined&judgeStatus=undefined&tags=580&title=

给定一个用单链表表示的整数，然后把这个整数加一。

示例1
输入：{1,2,3}
返回值：{1,2,4}

示例2
输入：{1,2,0}
返回值：{1,2,1}

示例3
输入：{9}
返回值：{1,0}
"""

# 思路1： 拿到最后1位， 进位处理， 有点类似整数加1
# 思路2： 遍历，拿到数据，塞到数组，然后处理，最后用链表串起来


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def plus_one(head):
    cur = head
    min_position = None
    # 找到不为9的最大下标
    while cur is not None:
        if cur.val != 9:
            min_position = cur
        cur = cur.next
    # 所有位都为9
    if min_position is None:
        # 初始化一个链表，头部为0
        min_position = ListNode(0, head)
        # 赋值给head（同步将head 头部加上0)
        head = min_position
    # 情况1： 所有位都不大于9(1,2,3,4)，下面的代码会将最后位加1，然后走向终点，for 循环进不去
    # 情况2： 所有位都为9(9,9,9,9)， 此时的minposition 和 head 都在最前面加上了0， 下面的代码会将最前的0 加上1， 变为1，for循环中，不停的将后面的9置为0
    # 情况3： 所有位不是都为9（9,9,8,9），此时的minposition 停留在不为9的最大下标 ， 下面的代码会将当前位的minposition 加上1， 如果能进入for循环中，不停的将后面的9置为0, 直到不能进入for循环，走向终点
    min_position.val += 1
    min_position = min_position.next
    while min_position is not None:
        min_position.val = 0
        min_position = min_position.next

    return head


# def main():
#     # test
#     node = ListNode(
#         1,
#         ListNode(
#             2,
#             ListNode(
#                 3,
#                 ListNode(
#                     4,
#                     ListNode(4, None),
#                 ),
#             ),
#         ),
#     )
#
#     result = plus_one(node)
#     print(result.val)
#     print(result.next.val)
#     print(result.next.next.val)
#     print(result.next.next.next.val)
