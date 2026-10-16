"""105. 从前序与中序遍历序列构造二叉树
题目：preorder + inorder 还原二叉树。
思路：前序首元素是根；中序中根的位置划分左右子树。
      用哈希表存中序的值→下标，避免每次线性查找。
复杂度：时间 O(n)，空间 O(n)。
"""

from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        idx = {v: i for i, v in enumerate(inorder)}
        self.pre_i = 0

        def build(l: int, r: int) -> Optional[TreeNode]:
            if l > r:
                return None
            root_val = preorder[self.pre_i]
            self.pre_i += 1
            root = TreeNode(root_val)
            m = idx[root_val]
            root.left = build(l, m - 1)       # 先建左：前序顺序 先根→左→右
            root.right = build(m + 1, r)
            return root

        return build(0, len(inorder) - 1)


if __name__ == "__main__":
    s = Solution()
    t = s.buildTree([3, 9, 20, 15, 7], [9, 3, 15, 20, 7])
    assert t.val == 3 and t.left.val == 9 and t.right.val == 20
    assert t.right.left.val == 15 and t.right.right.val == 7
    print("47 OK")
