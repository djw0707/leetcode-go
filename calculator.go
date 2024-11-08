package main

import (
	"fmt"
)

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func getIndex(s string) map[int]int {
	stack := make([]int, 0)
	indexMap := make(map[int]int)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			indexMap[top] = i
		} else {
			continue
		}
	}
	return indexMap
}

// 一个简单的加减法计算器
func calculate(s string, start int, end int) int {
	indexMap := getIndex(s)
	fmt.Println(indexMap)
	// 保存每一位带符号的数字
	stack := make([]int, 0)
	// 符号标志位
	pre := byte('+')
	// 记录当前数字
	num := 0
	for i := start; i <= end; i++ {
		if isDigit(s[i]) {
			// 优先计算数字
			num = num*10 + int(s[i]-'0')
		}
		if s[i] == '(' {
			num = calculate(s, i+1, indexMap[i]-1)
			i = indexMap[i]
			fmt.Println("calculate result: ", num, i)
		}
		// 最后判断i是否是最后一位是为了计算的时候不遗漏最后一个数
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' || i == end {
			if pre == '+' {
				stack = append(stack, num)
			} else if pre == '-' {
				stack = append(stack, -num)
			} else if pre == '*' {
				stack[len(stack)-1] = stack[len(stack)-1] * num
				fmt.Println("find *:", i, stack, num)
			} else if pre == '/' {
				stack[len(stack)-1] = stack[len(stack)-1] / num
			}
			pre = s[i]
			num = 0
		}
	}
	ans := 0
	for i := range stack {
		ans += stack[i]
	}
	return ans
}

func main() {
	s := "(1+(4+5+2)-3)+(6+8)"
	fmt.Println(calculate(s, 0, len(s)-1))
}
