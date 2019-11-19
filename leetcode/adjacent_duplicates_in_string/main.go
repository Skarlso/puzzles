package adjacent_duplicates_in_string

func removeDuplicates(s string, k int) string {
	sb := []byte(s)
	duplicateFound := false
main:
	for {
		var (
			removeFrom int
			removeTo   int
		)
		for i := 0; i < len(sb)-1; i++ {
			if sb[i] == sb[i+1] {
				removeFrom = i
				for j := i; j < i+k; j++ {
					if j < len(sb) && sb[j] == sb[i] {
						removeTo = j
					}
				}
				if removeTo-removeFrom == k-1 {
					duplicateFound = true
					sb = append(sb[:removeFrom], sb[removeTo+1:]...)
				}
			}
		}

		if !duplicateFound {
			break main
		}
		duplicateFound = false
	}
	return string(sb)
}
