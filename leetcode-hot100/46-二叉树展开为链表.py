"""114. 二叉树展开为链表
题目：按先序遍历顺序把树展开成"右指针链表"，原地。
思路：对每个节点：若有左子树，找到左子树的最右节点（前驱），
      把当前右子树接到前驱右下，左子树整体搬到右边，左置空。
复杂度：时间 O(n)（每节点最多被前驱访问一次），空间 O(1)。
"""

from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def flatten(self, root: Optional[TreeNode]) -> None:
        cur = root
        while cur:
            if cur.left:
                # 找左子树的最右节点（先序前驱）
                pred = cur.left
                while pred.right:
                    pred = pred.right
                pred.right = cur.right        # 右子树接前驱
                cur.right = cur.left          # 左子树搬右边
                cur.left = None
            cur = cur.right


def to_list(root):
    out = []
    while root:
        assert root.left is None
        out.append(root.val)
        root = root.right
    return out


if __name__ == "__main__":
    s = Solution()
    t = TreeNode(1, TreeNode(2, TreeNode(3), TreeNode(4)), TreeNode(5, None, TreeNode(6)))
    s.flatten(t)
    assert to_list(t) == [1, 2, 3, 4, 5, 6]
    print("46 OK")
