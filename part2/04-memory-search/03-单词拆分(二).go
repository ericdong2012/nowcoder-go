package _4_memory_search


/*
NC182 单词拆分(二)
https://www.nowcoder.com/practice/bb40e0aae84a4491a35f93b322516cb5?tpId=117&tqId=39349&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=1617&title=


描述
给定一个字符串 s 和一个字符串数组 dic ，在字符串 s 的任意位置添加任意多个空格后得到的字符串集合是给定字符串数组 dic 的子集（即拆分后的字符串集合中的所有字符串都在 dic 数组中），你可以以任意顺序 返回所有这些可能的拆分方案。

数据范围：字符串长度满足 1 \le n \le 20 \1≤n≤20  ，数组长度满足 1 \le n \le 10 \1≤n≤10  ，数组中字符串长度满足 1 \le n \le 20 \1≤n≤20
示例1
输入：
"nowcoder",["now","coder","no","wcoder"]
复制
返回值：
["no wcoder","now coder"]
复制
示例2
输入：
"nowcoder",["now","wcoder"]
复制
返回值：
[]
复制
示例3
输入：
"nowcoder",["nowcoder"]
复制
返回值：
["nowcoder"]
复制
示例4
输入：
"nownowcoder",["now","coder"]
复制
返回值：
["now now coder"]
复制
说明：
你可以重复使用 dic 数组中的字符串

*/