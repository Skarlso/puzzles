package longest_substring

import "bytes"

func lengthOfLongestSubstring(str string) int {
	var temp, longestString []byte
	for _, r := range str {
		if bytes.ContainsAny(temp, string(r)) {
			i := bytes.Index(temp, []byte{byte(r)})
			temp = append(temp[i+1:], byte(r))
		} else {
			temp = append(temp, byte(r))
		}
		if len(temp) > len(longestString) {
			longestString = temp
		}
	}
	return len(longestString)
}
