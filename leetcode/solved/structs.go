package gogo

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (p *ListNode) String() string {
	to_string := ""
	for p != nil {
		to_string += strconv.Itoa(p.Val) + " -> "
		p = p.Next
	}
	to_string += "nil"
	return to_string
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
