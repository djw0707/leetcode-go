// package mergedSet
package main

import "fmt"

type mergeSet struct {
	parent   []int
	capacity int
	rank     []int
}

func UnionFindInit(n int) *mergeSet {
	s := mergeSet{}
	s.capacity = n
	s.parent = make([]int, n)
	s.rank = make([]int, n)
	for i := 0; i < n; i++ {
		s.parent[i] = i
		s.rank[i] = 1
	}
	return &s
}

func (s *mergeSet) find(i int) int {
	if s.parent[i] == i {
		return i
	}
	return s.find(s.parent[i])
}

func (s *mergeSet) Connect(i int, j int) bool {
	if s.isConnected(i, j) {
		return false
	}
	s.parent[j] = i
	// r1 := s.rank[i]
	// r2 := s.rank[j]
	// if r1 <= r2 {
	// 	s.parent[i] = j
	// 	s.rank[j] += s.rank[i]
	// } else {
	// 	s.parent[j] = i
	// 	s.rank[i] += s.rank[j]
	// }
	return true
}

func (s *mergeSet) isConnected(i int, j int) bool {
	p1 := s.find(i)
	p2 := s.find(j)
	if p1 == p2 {
		return true
	}
	return false
}

func (s *mergeSet) Print() {

	for i := 0; i < s.capacity; i++ {
		if s.parent[i] != i {

		}
	}
	fmt.Println()
}

func main() {
	s := UnionFindInit(10)
	fmt.Println(s.Connect(1, 2))
	fmt.Println(s.Connect(2, 3))
	fmt.Println(s.Connect(3, 1))
	fmt.Println(s.Connect(4, 7))
	fmt.Println(s.Connect(5, 6))
	fmt.Println(s.Connect(4, 5))
}
