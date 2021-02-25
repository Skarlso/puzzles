package main

import "testing"

func TestMatch_1(t *testing.T) {
	type test struct {
		s       string
		pattern string
		want    bool
	}

	tests := []test{
		{s: "aa", pattern: "a", want: false},
		{s: "aa", pattern: "a*", want: true},
		{s: "ab", pattern: ".*", want: true},
		{s: "ab", pattern: "..", want: true},
		{s: "aab", pattern: "c*a*b", want: true},
		{s: "mississippi", pattern: "mis*is*p*.", want: false},
		{s: "ab", pattern: ".*c", want: false},
		{s: "aaa", pattern: "aaaa", want: false},
		{s: "a", pattern: "ab*", want: true},
		{s: "bbbba", pattern: ".*a*a", want: true},
	}
	for _, ts := range tests {
		got := isMatch(ts.s, ts.pattern)
		if got != ts.want {
			t.Fatal(got, ts.want, ts.s, ts.pattern)
		}
	}
}
