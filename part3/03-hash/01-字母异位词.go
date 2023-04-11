package _3_hash

/*
题⽬⼤意
给出 2 个字符串 s 和 t，如果 t 中的字⺟在 s 中都存在，输出 true，否则输出 false。

解题思路
这道题可以⽤打表的⽅式做。先把 s 中的每个字⺟都存在⼀个 26 个容量的数组⾥⾯，每个下标依次对
应 26 个字⺟。s 中每个字⺟都对应表中⼀个字⺟，每出现⼀次就加 1。然后再扫字符串 t，每出现⼀个
字⺟就在表⾥⾯减⼀。如果都出现了，最终表⾥⾯的值肯定都是 0 。最终判断表⾥⾯的值是否都是 0 即
可，有⾮ 0 的数都输出 false 。

*/
// 解法⼀
func isAnagram(s string, t string) bool {
	alphabet := make([]int, 26)
	sBytes := []byte(s)
	tBytes := []byte(t)
	if len(sBytes) != len(tBytes) {
		return false
	}
	for i := 0; i < len(sBytes); i++ {
		alphabet[sBytes[i]-'a']++
	}
	for i := 0; i < len(tBytes); i++ {
		alphabet[tBytes[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 {
			return false

		}
	}
	return true
}

// 解法⼆
func isAnagram1(s string, t string) bool {
	if s == "" && t == "" {
		return true
	}
	if s == "" || t == "" {
		return false
	}
	sBytes := []byte(s)
	tBytes := []byte(t)
	if len(sBytes) != len(tBytes) {
		return false
	}
	quickSortByte(sBytes, 0, len(sBytes)-1)
	quickSortByte(tBytes, 0, len(tBytes)-1)
	for i := 0; i < len(sBytes); i++ {
		if sBytes[i] != tBytes[i] {
			return false
		}
	}
	return true
}

func partitionByte(a []byte, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] > pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

func quickSortByte(a []byte, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionByte(a, lo, hi)
	quickSortByte(a, lo, p-1)
	quickSortByte(a, p+1, hi)
}
