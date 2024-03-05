package main


/*
https://www.nowcoder.com/practice/d1418aaa147a4cb394c3c3efc4302266?tpId=117&tqId=37844&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D196&difficulty=undefined&judgeStatus=undefined&tags=593&title=

NC87 丢棋子问题

一座大楼有 n+1 层，地面算作第0层，最高的一层为第 n 层。已知棋子从第0层掉落肯定不会摔碎，从第 i 层掉落可能会摔碎，也可能不会摔碎。

给定整数 n 作为楼层数，再给定整数 k 作为棋子数，返回如果想找到棋子不会摔碎的最高层数，即使在最差的情况下扔的最小次数。一次只能扔一个棋子。


示例1
输入：
10,1
复制
返回值：
10
复制
说明：
因为只有1棵棋子，所以不得不从第1层开始一直试到第10层，在最差的情况下，即第10层是不会摔坏的最高层，最少也要扔10次
示例2
输入：
3,2
复制
返回值：
2
复制
说明：
先在2层扔1棵棋子，如果碎了，试第1层，如果没碎，试第3层


*/