package main

import (
	"fmt"
	"testing"
)

func TestMath(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	// l1l := getList(l1)
	// l2l := getList(l2)
	// fmt.Println(l1l, l2l)
	l0 := addTwoNumbers(l1, l2)
	fmt.Println(getList(l0))
}

func TestZeros(t *testing.T) {
	v := addTwoNumbers(&ListNode{Val: 0}, &ListNode{Val: 0})
	fmt.Println(getList(v))
}

func TestOne(t *testing.T) {
	v := addTwoNumbers(&ListNode{Val: 1}, &ListNode{Val: 0})
	fmt.Println(getList(v))
}

func TestNotEqual(t *testing.T) {
	v := addTwoNumbers(&ListNode{Val: 9, Next: &ListNode{Val: 8}}, &ListNode{Val: 1})
	fmt.Println(getList(v))
}

func TestNotEqual2(t *testing.T) {
	v := addTwoNumbers(&ListNode{Val: 0}, &ListNode{Val: 1, Next: &ListNode{Val: 8}})
	fmt.Println(getList(v))
}
