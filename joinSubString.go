package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	smap := make(map[string]int)
	wmap := make(map[string]int)
	ans := make([]int, 0)
	consist := 0
	for i := 0; i < len(words); i++ {
		wmap[words[i]] += 1
	}
	l := len(words[0])
	// 起点从0到l-1，因为每次匹配可能不是从第一个字母开始的，如果l等于3 那从0,1,2位置都可能开始 如果从3开始
	// 和从0开始是一样的，因为滑动窗口每次移动l长度，如果0位置不满足条件下一个滑动位置就是l。然而这种情况
	// 无法覆盖到从1开始匹配的情况
	for k := 0; k < l; k++ {
		i, j := k, k
		smap = make(map[string]int)
		consist = 0
		// 每个位置都进行一次滑动窗口
		for j+l <= len(s) {
			// 这部分更新窗口右侧
			cur := s[j : j+l]
			j += l
			smap[cur] += 1
			_, ok := wmap[cur]
			if ok {
				// 记录一致的元素个数
				if smap[cur] == wmap[cur] {
					consist++
				}
			}
			// 这部分左移窗口的左侧 第一个条件保证没有字典外的单词
			// 第二个条件保证窗口内元素需要等于字典内元素
			for len(smap) > len(wmap) || smap[cur] > wmap[cur] {
				left := s[i : i+l]
				i += l
				_, ok := smap[left]
				if ok {
					// 这一步如何更新consist很重要，在缩小左窗口时只记录减小之后不一致的个数
					if smap[left] == wmap[left] {
						consist--
					}
					smap[left]--
					if smap[left] == 0 {
						delete(smap, left)
					}
				}
			}
			// 当一致时候记录位置
			if consist == len(wmap) {
				ans = append(ans, i)
			}
		}
	}
	return ans
}
func main() {
	fmt.Println(findSubstring("bcabbcaabbccacacbabccacaababcbb", []string{"c", "b", "a", "c", "a", "a", "a", "b", "c"}))
}
