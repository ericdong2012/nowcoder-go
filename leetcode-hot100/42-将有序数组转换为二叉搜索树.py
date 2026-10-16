"""108. 将有序数组转换为二叉搜索树
题目：升序数组转成高度平衡的 BST。
思路：取中点做根，左右两半递归——中点分割天然保证平衡。
复杂度：时间 O(n)，空间 O(log n)（递归栈）。
"""

from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        def build(l: int, r: int) -> Optional[TreeNode]:
            if l > r:
                return None
            mid = (l + r) // 2
            node = TreeNode(nums[mid])
            node.left = build(l, mid - 1)
            node.right = build(mid + 1, r)
            return node

        return build(0, len(nums) - 1)


def is_bst(node, lo=float("-inf"), hi=float("inf")) -> bool:
    if not node:
        return True
    return lo < node.val < hi and is_bst(node.left, lo, node.val) and is_bst(node.right, node.val, hi)


if __name__ == "__main__":
    s = Solution()
    t = s.sortedArrayToBST([-10, -3, 0, 5, 9])
    assert is_bst(t)
    assert t.val in (0, -3, 5)               # 平衡解不唯一，中点是 0
    assert t.val == 0
    print("42 OK")
