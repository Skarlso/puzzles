package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getList(l *ListNode) []int {
	n := l
	digits := make([]int, 0)
	for n != nil {
		digits = append(digits, n.Val)
		n = n.Next
	}
	return digits
}

func reverse(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func addList(l1, l2 []int) []int {
	var result []int
	var remains int
	finalIndex := 0
	for i := 0; i < len(l1); i++ {
		if i >= len(l2) {
			result = append(result, l1[i:]...)
			return result
		}
		sum := (l1[i] + l2[i]) + remains
		remains = sum / 10
		if remains > 0 {
			sum %= 10
			remains = 1
		}
		result = append(result, sum)
		finalIndex = i
	}
	if finalIndex < len(l2)-1 {
		result = append(result, l2[finalIndex+1:]...)
	}
	if remains != 0 {
		result = append(result, remains)
	}
	return result
}

func createLinkFromList(l []int) *ListNode {
	if len(l) < 1 {
		return nil
	}
	var l1 int
	l1, l = l[0], l[1:]
	start := &ListNode{Val: l1}
	if len(l) < 1 {
		return start
	}
	last := &ListNode{}
	start.Next = last
	for i, num := range l {
		last.Val = num
		if i < len(l)-1 {
			last.Next = &ListNode{}
			last = last.Next
		}
	}
	return start
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	v := addList(getList(l1), getList(l2))
	reverse(v)
	return createLinkFromList(v)
}
