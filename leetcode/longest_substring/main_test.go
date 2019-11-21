package longest_substring

import "testing"

func TestBasic(t *testing.T) {
	s := "abcabcbb"
	o := lengthOfLongestSubstring(s)
	if o != 3 {
		t.Fatal("o was: ", o)
	}
}

func Test2Basic(t *testing.T) {
	s := "bbbbb"
	o := lengthOfLongestSubstring(s)
	if o != 1 {
		t.Fatal("o was: ", o)
	}
}

func Test3Basic(t *testing.T) {
	s := "aab"
	o := lengthOfLongestSubstring(s)
	if o != 2 {
		t.Fatal("o was: ", o)
	}
}

func Test4Basic(t *testing.T) {
	s := "dvdf"
	o := lengthOfLongestSubstring(s)
	if o != 3 {
		t.Fatal("o was: ", o)
	}
}

func BenchmarkBasic(b *testing.B) {
	s := "abcabcbb"
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(s)
	}
}
