"""
https://www.nowcoder.com/practice/735a34ff4672498b95660f43b7fcd628

NC159 最小生成树

一个有 n 户人家的村庄，有 m 条路相互连接着。村里现在要修路，每条路都有一个成本价格，现在请你帮忙计算下，最少需要花费多少钱，就能让这 n 户人家连接起来。
为一个二维数组，每个元素是一个长度为 3 的一维数组，和 表示村庄和村庄有一条路，修这条路的成本价格为。
每户之间可能有多条道路连接，但不可能自己与自己相连。
数据范围: , ，
进阶： 时间复杂度，空间复杂度

示例1
输入
3,3,[[1,3,3],[1,2,1],[2,3,1]]
输出
2

示例2
输入
2,1,[[1,2,1]]
输出
1

NC159 最小生成树
"""

"""
NOTE: 原始 Go 文件只包含题目描述和 package 声明，没有实际算法实现。
这里补充 Kruskal 最小生成树算法实现（使用并查集）。
"""


class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n + 1))
        self.rank = [0] * (n + 1)

    def find(self, x):
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x, y):
        px, py = self.find(x), self.find(y)
        if px == py:
            return False
        if self.rank[px] < self.rank[py]:
            px, py = py, px
        self.parent[py] = px
        if self.rank[px] == self.rank[py]:
            self.rank[px] += 1
        return True


def mini_spanning_tree(n, m, cost):
    """
    Kruskal 算法：将所有边按权重排序，依次尝试加入边，用并查集判断是否形成环。
    n: 村庄数量 (节点编号 1..n)
    m: 道路数量
    cost: 二维数组，每个元素 [u, v, w] 表示村庄 u 和 v 之间修路成本 w
    返回: 最小生成树的总成本
    """
    # 按成本升序排列所有边
    cost.sort(key=lambda x: x[2])

    uf = UnionFind(n)
    total = 0
    edge_count = 0

    for u, v, w in cost:
        if uf.union(u, v):
            total += w
            edge_count += 1
            if edge_count == n - 1:
                break

    return total


if __name__ == "__main__":
    print(mini_spanning_tree(3, 3, [[1, 3, 3], [1, 2, 1], [2, 3, 1]]))  # 输出 2
    print(mini_spanning_tree(2, 1, [[1, 2, 1]]))  # 输出 1
