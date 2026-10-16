"""
NC6 二叉树中的最大路径和
https://www.nowcoder.com/practice/da785ea0f64b442488c125b441a4ba4a?tpId=117&tqId=37716&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117

描述
二叉树里面的路径被定义为:从该树的任意节点出发，经过父=>子 或者 子=>父的连接，达到任意节点的序列。
注意:
1. 同一个节点在一条二叉树路径里中最多出现一次
2. 一条路径至少包含一个节点，且不一定经过根节点

给定一个二叉树的根节点root，请你计算它的最大路径和

示例1
输入:
{1,2,3}
返回值:
6

示例2
输入:
{-20,8,20,#,#,15,6}
返回值:
41

示例3
输入:
{-2,#,-3}
返回值:
-2
"""


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


# 全局变量记录最大路径和
_res = -10000000


def max_path_sum(root):
    global _res
    _res = -10000000
    dfs(root)
    return _res


# 解法一
def dfs(root):
    global _res
    if root is None:
        return 0
    # 后序遍历
    # 分别求出左右子树的最大贡献度
    # 只有在最大贡献值大于 0 时，才会选取对应子节点
    left = max(dfs(root.left), 0)
    right = max(dfs(root.right), 0)
    # 根，根+左，根+右 三者的较大值
    cur_sum = max(root.val, root.val + right, root.val + left)
    # 较大值和根+左+右的比较
    cur_max = max(cur_sum, left + right + root.val)
    # 最大值
    _res = max(_res, cur_max)
    # 返回节点的最大贡献值 node.val + max(leftGain, rightGain)
    return cur_sum


# 解法2
def dfs2(root):
    global _res
    if root is None:
        return 0
    # 后序遍历
    left = max(dfs2(root.left), 0)
    right = max(dfs2(root.right), 0)
    # 最终的最大路径和 = 左子树最大深度 + 右子树最大深度 + 本节点，负值会被舍弃
    _res = max(_res, root.val + left + right)
    # 返回当前节点的最大路径
    num = max(left, right) + root.val
    return num


if __name__ == "__main__":
    node = TreeNode(
        val=3,
        left=TreeNode(val=9),
        right=TreeNode(
            val=20,
            left=TreeNode(val=15),
            right=TreeNode(val=7),
        ),
    )
    result = max_path_sum(node)
    print(result)
