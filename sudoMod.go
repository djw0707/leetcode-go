package main

import (
	"fmt"
	"strconv"
)

func isValidSudoku(board [][]byte) bool {
	n := len(board)
	m := len(board[0])
	row := make([]int, n)
	col := make([]int, m)
	sudoMap := [9][9]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == '.' {
				continue
			}
			bit, _ := strconv.Atoi(string(board[i][j]))
			if row[i]&1<<bit == 1<<bit {
				return false
			} else if col[j]&1<<bit == 1<<bit {
				return false
			}
			row[i] = 1<<bit | row[i]
			col[j] = 1<<bit | col[j]
			index := (i/3)*3 + j/3
			if sudoMap[index][bit-1] == 1 {
				return false
			}
			sudoMap[index][bit-1] = 1
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '6', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	}
	fmt.Println(isValidSudoku(board))
}

// map[97:3 103:1 109:1 110:1 114:1] map[97:3 103:1 109:1 110:1 114:1] 3
