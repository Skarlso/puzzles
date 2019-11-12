package main

import (
	"fmt"
	"strings"
)

func searchString(s, search string) (count int) {
	if len(search) == 1 {
		if strings.Contains(s, search) {
			count++
			return
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] == search[0] {
			count += searchString(string(s[i:]), string(search[1:]))
		}
	}
	return
}

func main() {
	s := "ADGGKJDBLKASDLFKSFKC"
	search := "ABC"
	fmt.Println(searchString(s, search))
}
