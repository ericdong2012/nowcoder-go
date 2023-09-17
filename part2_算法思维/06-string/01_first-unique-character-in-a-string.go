package main

import "strings"

/*
387. 字符串中的第一个唯一字符   找出字符串中第一个只出现一次的字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2


注意事项：您可以假定该字符串只包含小写字母。

https://www.nowcoder.com/practice/e896d0f82f1246a3aa7b232ce38029d4

*/

func firstUniqChar(s string) int {
    for _, v := range "abcdefghijkllmnopqrstuvwxyz" {
       index := strings.Index(s, string(v))
       lastIndex := strings.LastIndex(s, string(v))
       if index != -1 && lastIndex != -1 && index == lastIndex {
           return index
       }
    }
    return -1

    //// 使用哈希表记录每个字符出现的次数
    //count := make(map[rune]int)
    //for _, ch := range s {
    //    count[ch]++
    //}
    //// 第一遍扫描完毕后，再进行一遍扫描，找到第一个只出现一次的字符
    //for i, ch := range s {
    //    if count[ch] == 1 {
    //        return i
    //    }
    //}
    //// 如果找不到只出现一次的字符，则返回 -1
    //return -1
}

