package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	c := searchString("rabbbit", "rabbit")
	if c != 3 {
		t.Fatal(c)
	}
}

func TestMain2(t *testing.T) {
	c := searchString("ABCBABC", "ABC")
	if c != 5 {
		t.Fatal(c)
	}
}

func TestMain3(t *testing.T) {
	c := searchString("babgbag", "bag")
	if c != 5 {
		t.Fatal(c)
	}
}

func xTestBench(t *testing.T) {
	s := "aaddbacbbcabdbceaeeaacbabbbaccacaacbabeddbedcdceceeabccabdadccceebcbcbecdbacedcecdeadbaebceaedaaecbbebdecabbddccebaccabaaabbabceddceecadcccdceabbcdadbbadebdadeacbaddccdeebcddaebbcbedbd"
	search := "ccdeddeabb"
	searchString(s, search)
}
