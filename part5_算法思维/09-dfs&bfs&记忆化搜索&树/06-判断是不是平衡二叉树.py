"""
BM36 判断是不是平衡二叉树
https://www.nowcoder.com/practice/8b3b95850edb4115918ecebdf1b4d222?tpId=295&tqId=23250&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
输入一棵节点数为 n 二叉树，判断该二叉树是否是平衡二叉树。
平衡二叉树（Balanced Binary Tree）：
它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是平衡二叉树。

示例1
输入：
{1,2,3,4,5,6,7}
返回值：
true

示例2
输入：
{}
返回值：
true
"""

# 思路1：
# 1.判断以根结点的树是否为二叉平衡树。求出左右子树的高度，判断它们的高度差是否超过了1。
# 2.递归判断根的左子树是否为平衡二叉树
# 3.递归判断根的右子树是否为平衡二叉树
# 注意：空树也是平衡二叉树


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def is_balanced_solution(p_root):
    """write code here"""
    if p_root is None:
        return True
    # {1,2,3,4,#,#,5,6}  防止单侧层级不满足要求
    if not is_balanced_solution(p_root.left):
        return False
    if not is_balanced_solution(p_root.right):
        return False
    # 判断左右子树的高度差 是否小于等于1
    return abs(get_max_depth(p_root.left) - get_max_depth(p_root.right)) <= 1


# 获取左子树或者右子树的深度
# 不是通过算树的深度，而是直接迭代
def get_max_depth(p_root):
    if p_root is None:
        return 0
    h1 = get_max_depth(p_root.left)
    h2 = get_max_depth(p_root.right)
    # 加上根节点
    return 1 + max(h1, h2)


if __name__ == "__main__":
    node = TreeNode(
        val=3,
        left=TreeNode(val=1),
        right=TreeNode(
            val=2,
            left=TreeNode(val=4),
            right=TreeNode(val=5),
        ),
    )
    result = is_balanced_solution(node)
    print(result)
