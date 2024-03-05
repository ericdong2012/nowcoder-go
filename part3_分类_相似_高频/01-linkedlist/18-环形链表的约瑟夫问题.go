package main

/*
https://www.nowcoder.com/practice/41c399fdb6004b31a6cbb047c641ed8a?tpId=117&tqId=37812&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D196&difficulty=undefined&judgeStatus=undefined&tags=580&title=

NC132 环形链表的约瑟夫问题

编号为 1 到 n 的 n 个人围成一圈。从编号为 1 的人开始报数，报到 m 的人离开。
下一个人继续从 1 开始报数。
n-1 轮结束以后，只剩下一个人，问最后留下的这个人编号是多少？


示例1
输入：
5,2
复制
返回值：
3
说明：
开始5个人 1，2，3，4，5 ，从1开始报数，1->1，2->2编号为2的人离开
1，3，4，5，从3开始报数，3->1，4->2编号为4的人离开
1，3，5，从5开始报数，5->1，1->2编号为1的人离开
3，5，从3开始报数，3->1，5->2编号为5的人离开
最后留下人的编号是3


示例2
输入：
1,1
复制
返回值：
1


*/
