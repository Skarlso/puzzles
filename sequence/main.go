package main

import (
	"fmt"
	"strings"
)

// Chop up the string as we search it and chop up the search string as well
// once the first character matched. If the first character matched
// we only have to further look for the rest of them to match.
// Once there is only one character left in the search string we can just
// see if it occurs in the string. If it does we know our sequence is complete
// so we can increase the count by one.
// Once that happens, the recurisve call unvineds and continues with the next
// character in the tail of the string until the search string is empty.
// 58 / 63 test cases passed. on leetcode. Time limits are not passed.
func searchString(s, search string) (count int) {
	if len(search) == 1 {
		if c := strings.Count(s, search); c > -1 {
			count += c
			return
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] == search[0] {
			count += searchString(s[i+1:], search[1:])
		}
	}
	return
}

func main() {
	s := "babgbag"
	search := "bag"
	// fmt.Println(s, search)
	count := 0
	for i := 0; i < len(s); i++ {
		sch := search
		s2 := s
		if s2[i] == sch[0] {
			sch = sch[1:]
			s2 = s2[i+1:]
			for len(sch) > 1 {

			}
			// for j := 0; j < len(sch); j++ {
			// 	if
			// }
		}
	}
	fmt.Println(count)
}
