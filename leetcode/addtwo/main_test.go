package main

import (
	"reflect"
	"testing"
)

func TestMath(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	l0 := addTwoNumbers(l1, l2)
	got := getList(l0)
	if !reflect.DeepEqual(got, []int{7, 0, 8}) {
		t.Fatal("got: ", got)
	}
}

func TestZeros(t *testing.T) {
	v := addTwoNumbers(&ListNode{Val: 0}, &ListNode{Val: 0})
	got := getList(v)
	if !reflect.DeepEqual(got, []int{0}) {
		t.Fatal("got: ", got)
	}
}

func TestOne(t *testing.T) {
	v := addTwoNumbers(&ListNode{Val: 1}, &ListNode{Val: 0})
	got := getList(v)
	if !reflect.DeepEqual(got, []int{1}) {
		t.Fatal("got: ", got)
	}
}

func TestNotEqual(t *testing.T) {
	v := addTwoNumbers(&ListNode{Val: 9, Next: &ListNode{Val: 8}}, &ListNode{Val: 1})
	got := getList(v)
	if !reflect.DeepEqual(got, []int{0, 9}) {
		t.Fatal("got: ", got)
	}
}

func Test3NotEqual(t *testing.T) {
	v := addTwoNumbers(&ListNode{Val: 1, Next: &ListNode{Val: 8}}, &ListNode{Val: 0})
	got := getList(v)
	if !reflect.DeepEqual(got, []int{1, 8}) {
		t.Fatal("got: ", got)
	}
}

func Test2NotEqual(t *testing.T) {
	v := addTwoNumbers(&ListNode{Val: 0}, &ListNode{Val: 1, Next: &ListNode{Val: 8}})
	got := getList(v)
	if !reflect.DeepEqual(got, []int{1, 8}) {
		t.Fatal("got: ", got)
	}
}

func Test4NotEqual(t *testing.T) {
	v := addTwoNumbers(&ListNode{Val: 1}, &ListNode{Val: 9, Next: &ListNode{Val: 9}})
	got := getList(v)
	if !reflect.DeepEqual(got, []int{0, 0, 1}) {
		t.Fatal("got: ", got)
	}
}
