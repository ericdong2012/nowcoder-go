"""437. 路径总和 III（前缀和 + 回溯）
题目：统计和为 targetSum 的"向下"路径条数（节点值可负，路径不必从根/到叶）。
思路：根到当前节点的前缀和 cur；路径和 == target ⇔ 存在祖先前缀和
      等于 cur - target。哈希表计数祖先前缀和，回溯时撤销。
复杂度：时间 O(n)，空间 O(n)。
"""

from typing import Optional
from collections import defaultdict


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:
        prefix = defaultdict(int)
        prefix[0] = 1                       # 空前缀：从根开始的路径
        self.count = 0

        def dfs(node, cur) -> None:
            if not node:
                return
            cur += node.val
            self.count += prefix[cur - targetSum]
            prefix[cur] += 1
            dfs(node.left, cur)
            dfs(node.right, cur)
            prefix[cur] -= 1                # 回溯撤销：祖先前缀只对本子树可见

        dfs(root, 0)
        return self.count


if __name__ == "__main__":
    s = Solution()
    t = TreeNode(10, TreeNode(5, TreeNode(3, TreeNode(3), TreeNode(-2)),
                              TreeNode(2, None, TreeNode(1))),
                 TreeNode(-3, None, TreeNode(11)))
    assert s.pathSum(t, 8) == 3
    print("48 OK")
