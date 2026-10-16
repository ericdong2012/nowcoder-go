"""199. 二叉树的右视图
题目：从右侧看每层能看到的节点值。
思路：BFS 每层取最后一个；或 DFS 根→右→左，首次到达的深度记录答案。
复杂度：时间 O(n)，空间 O(w)。
"""

from typing import List, Optional
from collections import deque


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        if not root:
            return []
        res, q = [], deque([root])
        while q:
            res.append(q[-1].val)             # 本层最后一个
            for _ in range(len(q)):
                node = q.popleft()
                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)
        return res


if __name__ == "__main__":
    s = Solution()
    t = TreeNode(1, TreeNode(2, None, TreeNode(5)), TreeNode(3, None, TreeNode(4)))
    assert s.rightSideView(t) == [1, 3, 4]
    # 左子树更深：右视图看到的可能是左子树节点
    t2 = TreeNode(1, TreeNode(2, TreeNode(4)), TreeNode(3))
    assert s.rightSideView(t2) == [1, 3, 4]
    print("45 OK")
