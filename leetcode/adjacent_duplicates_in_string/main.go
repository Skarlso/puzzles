package adjacent_duplicates_in_string

/*
Runtime: 264 ms, faster than 10.64% of Go online submissions for Remove All Adjacent Duplicates in String II.
Memory Usage: 3.4 MB, less than 100.00% of Go online submissions for Remove All Adjacent Duplicates in String II.
*/
func removeDuplicates(s string, k int) string {
	sb := []byte(s)
	for {
		var (
			removeFrom     int
			removeTo       int
			duplicateFound bool
		)
		for i := 0; i < len(sb)-1; i++ {
			if sb[i] == sb[i+1] {
				removeFrom = i
				for j := i; j < i+k && j < len(sb); j++ {
					if sb[j] == sb[i] {
						removeTo = j
					} else {
						break
					}
				}
				if removeTo-removeFrom == k-1 {
					duplicateFound = true
					sb = append(sb[:removeFrom], sb[removeTo+1:]...)
				}
			}
		}

		if !duplicateFound {
			break
		}
	}
	return string(sb)
}
