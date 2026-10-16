"""102. 二叉树的层序遍历（BFS）
题目：按层输出节点值（二维数组）。
思路：队列 BFS；每轮记录当前队列长度 = 该层节点数，逐层弹出。
复杂度：时间 O(n)，空间 O(w)（最宽层）。
"""

from typing import List, Optional
from collections import deque


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if not root:
            return []
        res, q = [], deque([root])
        while q:
            level = []
            for _ in range(len(q)):          # 只处理本层
                node = q.popleft()
                level.append(node.val)
                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)
            res.append(level)
        return res


if __name__ == "__main__":
    s = Solution()
    t = TreeNode(3, TreeNode(9), TreeNode(20, TreeNode(15), TreeNode(7)))
    assert s.levelOrder(t) == [[3], [9, 20], [15, 7]]
    assert s.levelOrder(None) == []
    print("41 OK")
