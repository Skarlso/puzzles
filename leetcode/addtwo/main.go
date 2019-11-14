package main

// ListNode is a list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getList(l *ListNode) []int {
	digits := make([]int, 0)
	for l != nil {
		digits = append(digits, l.Val)
		l = l.Next
	}
	return digits
}

func addList(l1, l2 []int) []int {
	var (
		result  []int
		remains int
		i       int
	)
	for {
		if i >= len(l1) && i >= len(l2) {
			break
		}
		var sum int
		if i >= len(l2) {
			sum = l1[i] + remains
		} else if i >= len(l1) {
			sum = l2[i] + remains
		} else {
			sum = (l1[i] + l2[i]) + remains
		}
		remains = sum / 10
		if remains > 0 {
			sum %= 10
			remains = 1
		}
		result = append(result, sum)
		i++
	}
	if remains != 0 {
		result = append(result, remains)
	}
	return result
}

func createLinkFromList(l []int) *ListNode {
	start := &ListNode{Val: l[0]}
	if len(l) < 2 {
		return start
	}
	last := &ListNode{}
	start.Next = last
	for i := 1; i < len(l)-1; i++ {
		last.Val = l[i]
		last.Next = &ListNode{}
		last = last.Next
	}
	last.Val = l[len(l)-1]
	return start
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * ACCEPTED.
 * Runtime: 8 ms, faster than 90.20% of Go online submissions for Add Two Numbers.
 * Memory Usage: 6.1 MB, less than 7.32% of Go online submissions for Add Two Numbers.
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return createLinkFromList(addList(getList(l1), getList(l2)))
}
