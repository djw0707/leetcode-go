package main

// package miniSubstringLength

import (
	"fmt"
)

// go语言不支持两个切片或者map进行比较 因此需要一个变量来记录两个map有多少个key的值一致

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	// 记录左右边界
	ansL, ansR := -1, 1<<32-1
	// 记录一致的字母数
	consist := 0
	smap := make(map[byte]int, 0)
	tmap := make(map[byte]int, 0)
	for i := range t {
		tmap[t[i]]++
	}
	i, j := 0, 0
	// consist 记录smap中多少个key与tmap一致
	for j < len(s) {
		// 第一步扩大右窗口
		c := s[j]
		j++
		if cnt, ok := tmap[c]; ok {
			// 如果smap某个字母数多了也无所谓，后面缩小左端点会减
			smap[c]++
			if smap[c] == cnt {
				consist += 1
			}
		}
		// 第二步缩小左窗口 缩小的依据就是smap中key与tmap一致的个数，
		for consist == len(tmap) {
			// 先计算最小长度并更新
			if j-i < ansR-ansL {
				ansL = i
				ansR = j
			}
			c := s[i]
			i++
			// 缩小左窗口并从smap中除去该字母
			if _, ok := tmap[c]; ok {
				// 在除去之前做判断，如果smap等于tmap，说明移除之后不一致，valid应该减1，先减valid再更新smap
				if smap[c] == tmap[c] {
					consist--
				}
				smap[c]--
			}
		}
	}
	if ansL == -1 || ansR == -1 {
		return ""
	}
	return s[ansL:ansR]
}
func main() {
	fmt.Println(minWindow("AB", "B"))
}
