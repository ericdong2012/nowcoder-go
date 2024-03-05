package main

/*
https://www.nowcoder.com/practice/c9d44b73dc7c4dbfa4272224b1f9b42c?tpId=117&tqId=39450&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D196&difficulty=undefined&judgeStatus=undefined&tags=593&title=

NC225 三角形最小路径和

给定一个正三角形数组，自顶到底分别有 1，2，3，4，5...，n 个元素，找出自顶向下的最小路径和。

每一步只能移动到下一行的相邻节点上，相邻节点指下行种下标与之相同或下标加一的两个节点。



示例1
输入：
[[10]]
复制
返回值：
10
复制
示例2
输入：
[[2],[3,4],[6,5,7],[4,1,8,3]]
复制
返回值：
11
复制
说明：
最小路径是 2 ， 3 ，5 ， 1
示例3
输入：
[[1],[-1000,0]]
复制
返回值：
-999

*/