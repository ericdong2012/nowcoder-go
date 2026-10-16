"""124. 二叉树中的最大路径和
题目：任意节点到任意节点（向下路径）最大和，可不过根。
思路：类似 543——DFS 返回"以该节点为端点向下延伸"的最大贡献（负则取 0），
      全局最大 = 节点值 + 左贡献 + 右贡献（路径拐弯处）。
复杂度：时间 O(n)，空间 O(h)。
"""

from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        self.best = float("-inf")

        def gain(node) -> int:
            if not node:
                return 0
            l = max(gain(node.left), 0)     # 贡献为负不如不接
            r = max(gain(node.right), 0)
            self.best = max(self.best, node.val + l + r)   # 拐弯路径
            return node.val + max(l, r)     # 向上只能选一边

        gain(root)
        return self.best


if __name__ == "__main__":
    s = Solution()
    t = TreeNode(-10, TreeNode(9), TreeNode(20, TreeNode(15), TreeNode(7)))
    assert s.maxPathSum(t) == 42            # 15 + 20 + 7
    assert s.maxPathSum(TreeNode(-3)) == -3
    t2 = TreeNode(1, TreeNode(2), TreeNode(3))
    assert s.maxPathSum(t2) == 6
    print("50 OK")
