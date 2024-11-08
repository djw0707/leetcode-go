package main

import "strings"

func transIntToRome(num int) byte {
	switch num {
	case 1000:
		return 'M'
	case 500:
		return 'D'
	case 100:
		return 'C'
	case 50:
		return 'L'
	case 10:
		return 'X'
	case 5:
		return 'V'
	case 1:
		return 'I'
	default:
		return ' '
	}
}

func transBit(num int, cur int) string {
	if num > 0 && num < 4 {
		return strings.Repeat(string(transIntToRome(cur)), num)
	} else if num == 4 {
		return string(transIntToRome(cur)) + string(transIntToRome(5*cur))
	} else if num == 5 {
		return string(transIntToRome(5 * cur))
	} else if num > 5 && num < 9 {
		return string(transIntToRome(5*cur)) + strings.Repeat(string(transIntToRome(cur)), num-5)
	} else if num == 9 {
		return string(transIntToRome(cur)) + string(transIntToRome(10*cur))
	}
	return ""
}

func intToRoman(num int) string {
	ans := ""
	cur := 1000
	for cur >= 1 {
		bit := num / cur
		ans = ans + transBit(bit, cur)
		num = num % cur
		cur = cur / 10
	}
	return ans
}
