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
