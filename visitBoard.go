package main

import (
	"fmt"
)

func getBoardIndex(index int, last int) (int, int) {
	r := (index - 1) / last
	c := (index - 1) % last
	if r%2 == 1 {
		c = last - c - 1
	}
	r = last - r - 1
	return r, c
}

func snakesAndLadders(board [][]int) int {
	last := len(board) * len(board)
	queue := make([]int, 0)
	queue = append(queue, 1)
	isVisited := make([]bool, last+1)
	cnt := 0
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			for j := queue[i] + 1; j <= queue[i]+6 && j <= last; j++ {
				if j == last {
					return cnt + 1
				}
				r, c := getBoardIndex(j, len(board))
				if board[r][c] != -1 {
					queue = append(queue, board[r][c])
					isVisited[board[r][c]] = true
				}
				queue = append(queue, j)
				isVisited[j] = true
			}
		}
		cnt++
		queue = queue[l:]
	}
	return cnt
}

func main() {
	cnt := snakesAndLadders([][]int{{-1, -1, -1, 46, 47, -1, -1, -1}, {51, -1, -1, 63, -1, 31, 21, -1}, {-1, -1, 26, -1, -1, 38, -1, -1}, {-1, -1, 11, -1, 14, 23, 56, 57}, {11, -1, -1, -1, 49, 36, -1, 48}, {-1, -1, -1, 33, 56, -1, 57, 21}, {-1, -1, -1, -1, -1, -1, 2, -1}, {-1, -1, -1, 8, 3, -1, 6, 56}})
	fmt.Println(cnt)
}

//{{-1,-1,-1,46,47,-1,-1,-1},{51,-1,-1,63,-1,31,21,-1},{-1,-1,26,-1,-1,38,-1,-1},{-1,-1,11,-1,14,23,56,57},{11,-1,-1,-1,49,36,-1,48},{-1,-1,-1,33,56,-1,57,21},{-1,-1,-1,-1,-1,-1,2,-1},{-1,-1,-1,8,3,-1,6,56}}
