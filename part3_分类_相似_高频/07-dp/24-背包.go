package main

/*

https://www.nowcoder.com/practice/2820ea076d144b30806e72de5e5d4bbf?tpId=117&tqId=37856&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D196&difficulty=undefined&judgeStatus=undefined&tags=593&title=


NC145 01背包

已知一个背包最多能容纳体积之和为v的物品

现有 n 个物品，第 i 个物品的体积为 vi , 重量为 wi

求当前背包最多能装多大重量的物品?


示例1
输入：
10,2,[[1,3],[10,4]]
复制
返回值：
4
复制
说明：
第一个物品的体积为1，重量为3，第二个物品的体积为10，重量为4。只取第二个物品可以达到最优方案，取物重量为4
示例2
输入：
10,2,[[1,3],[9,8]]
复制
返回值：
11
复制
说明：
两个物品体积之和等于背包能装的体积，所以两个物品都取是最优方案

*/