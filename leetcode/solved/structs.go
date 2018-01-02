package gogo

type ListNode struct {
	Val  int
	Next *ListNode
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
