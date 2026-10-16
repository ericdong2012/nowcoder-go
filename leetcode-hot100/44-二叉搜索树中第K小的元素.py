"""230. 二叉搜索树中第 K 小的元素
题目：BST 中第 k 小的值。
思路：中序遍历天然有序——迭代版数到第 k 个即停（可提前终止）。
复杂度：时间 O(H + k)，空间 O(H)。
"""

from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        stack, cur = [], root
        count = 0
        while cur or stack:
            while cur:
                stack.append(cur)
                cur = cur.left
            cur = stack.pop()
            count += 1
            if count == k:
                return cur.val
            cur = cur.right
        raise ValueError("k 超出节点数")


def build():
    #     5
    #    / \
    #   3   6
    #  / \
    # 2   4
    return TreeNode(5, TreeNode(3, TreeNode(2), TreeNode(4)), TreeNode(6))


if __name__ == "__main__":
    s = Solution()
    t = build()
    assert s.kthSmallest(t, 3) == 4
    assert s.kthSmallest(t, 1) == 2
    assert s.kthSmallest(t, 5) == 6
    print("44 OK")
