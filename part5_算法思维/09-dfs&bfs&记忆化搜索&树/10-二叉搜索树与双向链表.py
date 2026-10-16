"""
BM30 二叉搜索树与双向链表
https://www.nowcoder.com/practice/947f6eb80d944a84850b0538bf0ec3a5?tpId=295&tqId=23253&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。如下图所示
要求：空间复杂度 O(1)（即在原树上操作），时间复杂度 O(n)

注意:
1.要求不能创建任何新的结点，只能调整树中结点指针的指向。当转化完成以后，
树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继
2.返回链表中的第一个节点的指针
3.函数返回的 TreeNode，有左右指针，其实可以看成一个双向链表的数据结构
4.你不用输出双向链表，程序会根据你的返回值自动打印输出

示例1
输入：{10,6,14,4,8,12,16}
返回值：From left to right are:4,6,8,10,12,14,16;From right to left are:16,14,12,10,8,6,4;
"""

# 将中序遍历的结果用数组存储下来，得到的数组是从小到大顺序的。最后将数组中的结点依次连接即可。


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def convert(root):
    """write code here"""
    if root is None:
        return None
    state = {"prev": None, "head": None}
    inorder(root, state)
    return state["head"]


def inorder(cur, state):
    """中序遍历，原地构建双向链表。

    使用 state dict 持有 prev 和 head，替代 Go 的包级变量。
    """
    if cur is None:
        return
    inorder(cur.left, state)
    # 如果前一个节点是空，说明找到了最左边，对 head 初始化
    if state["prev"] is None:
        state["head"] = cur
    else:
        # 构建双向链表
        state["prev"].right = cur
        cur.left = state["prev"]
    # 往下走了一步，类似++
    state["prev"] = cur
    inorder(cur.right, state)


if __name__ == "__main__":
    # 示例1 的 BST：{10,6,14,4,8,12,16}，转换后应打印有序序列 4,6,8,10,12,14,16
    node = TreeNode(
        val=10,
        left=TreeNode(
            val=6,
            left=TreeNode(val=4),
            right=TreeNode(val=8),
        ),
        right=TreeNode(
            val=14,
            left=TreeNode(val=12),
            right=TreeNode(val=16),
        ),
    )
    result = convert(node)
    # 中序遍历的头节点没有前驱(left 为 None)，沿 right 打印整个双向链表
    # (原 Go 版本访问 result.Left.Val 会在头节点处 nil 崩溃，这里改为正向遍历打印)
    cur = result
    while cur is not None:
        print(cur.val)
        cur = cur.right
