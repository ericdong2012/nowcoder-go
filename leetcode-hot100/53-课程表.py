"""207. 课程表（拓扑排序 / 判有向无环）
题目：numCourses 门课，prerequisites [a, b] 表示先修 b 才能修 a，判断能否完成。
思路：Kahn 算法——入度表 + 队列；入度 0 的先出队，出队后削减邻居入度；
      最终出队数 == 总课程数则无环。
复杂度：时间 O(V+E)，空间 O(V+E)。
"""

from typing import List
from collections import deque, defaultdict


class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        graph = defaultdict(list)
        indegree = [0] * numCourses
        for a, b in prerequisites:
            graph[b].append(a)
            indegree[a] += 1

        q = deque(i for i in range(numCourses) if indegree[i] == 0)
        taken = 0
        while q:
            c = q.popleft()
            taken += 1
            for nxt in graph[c]:
                indegree[nxt] -= 1
                if indegree[nxt] == 0:
                    q.append(nxt)
        return taken == numCourses


if __name__ == "__main__":
    s = Solution()
    assert s.canFinish(2, [[1, 0]]) is True
    assert s.canFinish(2, [[1, 0], [0, 1]]) is False
    assert s.canFinish(4, [[1, 0], [2, 0], [3, 1]]) is True
    print("53 OK")
