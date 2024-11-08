package main

import (
	"fmt"
)

type node struct {
	s   string
	val float64
}

func printMap(m map[string][]*node) {
	for k, v := range m {
		fmt.Println(k)
		for i := 0; i < len(v); i++ {
			fmt.Println(v[i].s, v[i].val)
		}
	}
}

func findAns(querie []string, m map[string][]*node) float64 {
	_, ok1 := m[querie[0]]
	_, ok2 := m[querie[1]]
	if !ok1 || !ok2 {
		return float64(-1)
	}
	que := make([]*node, 0)
	que = append(que, &node{s: querie[0], val: 1})
	visited := make(map[string]bool)
	for len(que) > 0 {
		first := que[0]
		arr := m[first.s]
		visited[first.s] = true
		if first.s == querie[1] {
			return first.val
		}
		que = que[1:]
		for i := 0; i < len(arr); i++ {
			if visited[arr[i].s] {
				continue
			}
			newNode := &node{s: arr[i].s, val: first.val * arr[i].val}
			que = append(que, newNode)
		}
	}
	return float64(-1)
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	noDegreeMap := make(map[string][]*node)
	ans := make([]float64, 0)
	for i := 0; i < len(equations); i++ {
		e := equations[i]
		if _, ok := noDegreeMap[e[0]]; ok {
			noDegreeMap[e[0]] = append(noDegreeMap[e[0]], &node{
				s:   e[1],
				val: values[i],
			})
		} else {
			noDegreeMap[e[0]] = []*node{&node{
				s:   e[1],
				val: values[i],
			}}
		}
		if _, ok := noDegreeMap[e[1]]; ok {
			noDegreeMap[e[1]] = append(noDegreeMap[e[1]], &node{
				s:   e[0],
				val: 1 / values[i],
			})
		} else {
			noDegreeMap[e[1]] = []*node{&node{
				s:   e[0],
				val: 1 / values[i],
			}}
		}
	}
	printMap(noDegreeMap)
	for i := 0; i < len(queries); i++ {
		ans = append(ans, findAns(queries[i], noDegreeMap))
	}
	printMap(noDegreeMap)
	return ans
}

func main() {
	fmt.Println(calcEquation([][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}, []float64{3.0, 4.0, 5.0, 6.0}, [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}}))
}
