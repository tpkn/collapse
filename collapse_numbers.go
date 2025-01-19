package collapse

// isEqualToPrevious checks the previous number is sequential to the current one
func isEqualToPrevious(s []int, i int) bool {
	return i-1 >= 0 && (s[i] == s[i-1] || s[i] == s[i-1]+1)
}

// isEqualToFollowing checks the next number is sequential to the current one
func isEqualToFollowing(s []int, i int) bool {
	return i+1 < len(s) && (s[i] == s[i+1] || s[i] == s[i+1]-1)
}

// isSequenceStart checks if the current number is the first in the sequence
func isSequenceStart(s []int, i int) bool {
	return i != len(s)-1 && !isSequenceMiddle(s, i) && isEqualToFollowing(s, i)
}

// isSequenceMiddle checks if the current number is somewhere in between of the sequence
func isSequenceMiddle(s []int, i int) bool {
	return isEqualToPrevious(s, i) && isEqualToFollowing(s, i)
}

// isSequenceEnd checks if the current number is the last in the sequence
func isSequenceEnd(s []int, i int) bool {
	return isEqualToPrevious(s, i) && !isEqualToFollowing(s, i)
}

// isNotInSequence checks if the current number is NOT in the sequence
func isNotInSequence(s []int, i int) bool {
	return !isSequenceStart(s, i) && !isSequenceMiddle(s, i) && !isSequenceEnd(s, i)
}

// Numbers collapses a sorted slice of numbers into {{start, end}, {...}} format
func Numbers(list []int) [][]int {
	var result [][]int
	var list_len = len(list)
	
	switch list_len {
	case 0:
		result = [][]int{}
	case 1:
		result = [][]int{{list[0], list[0]}}
	case 2:
		if list[0] == list[1] || list[0] == list[1]-1 {
			result = [][]int{{list[0], list[1]}}
		} else {
			result = [][]int{{list[0], list[0]}, {list[1], list[1]}}
		}
	default:
		var start, end int
		for i := 0; i < list_len; i++ {
			if isSequenceStart(list, i) {
				start = list[i]
				continue
			}
			if isSequenceMiddle(list, i) {
				continue
			}
			if isSequenceEnd(list, i) {
				end = list[i]
			}
			if isNotInSequence(list, i) {
				start = list[i]
				end = list[i]
			}
			result = append(result, []int{start, end})
		}
	}
	
	return result
}
