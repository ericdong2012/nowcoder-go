package _6_monotonous_stack

/*
NC237 最大矩形
https://www.nowcoder.com/practice/5720efc1bdff4ca3a7dad37ca012cb60?tpId=117&tqId=39474&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=1579&title=


描述
给定一个仅包含 0 和 1 ，大小为 n*m 的二维二进制矩阵，找出仅包含 1 的最大矩形并返回面积。

数据范围：1 \le n,m \le 200 \1≤n,m≤200  保证输入的矩形中仅含有 0 和 1

例如输入[[1,0,1,0,0],[1,1,1,1,0],[1,1,1,1,1],[1,0,0,1,0]]时，对应的输出为8，所形成的最大矩形如下图所示：

示例1
输入：
[[1]]
复制
返回值：
1
复制
示例2
输入：
[[1,1],[0,1]]
复制
返回值：
2
复制
说明：
第一列的两个 1 组成一个矩形
示例3
输入：
[[0,0],[0,0]]
复制
返回值：
0
复制
示例4
输入：
[[1,0,1,0,0],[1,1,1,1,0],[1,1,1,1,1],[1,0,0,1,0]]
复制
返回值：
8

*/
