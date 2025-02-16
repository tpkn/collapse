package collapse

// isEqualToPrevious checks the previous number is sequential to the current one
func isEqualToPrevious(s []int, i int, step int) bool {
	return i-1 >= 0 && (s[i] == s[i-1] || s[i] == s[i-1]+step)
}

// isEqualToFollowing checks the next number is sequential to the current one
func isEqualToFollowing(s []int, i int, step int) bool {
	return i+1 < len(s) && (s[i] == s[i+1] || s[i] == s[i+1]-step)
}

// isSequenceStart checks if the current number is the first in the sequence
func isSequenceStart(s []int, i int, step int) bool {
	return i != len(s)-1 && !isSequenceMiddle(s, i, step) && isEqualToFollowing(s, i, step)
}

// isSequenceMiddle checks if the current number is somewhere in between of the sequence
func isSequenceMiddle(s []int, i int, step int) bool {
	return isEqualToPrevious(s, i, step) && isEqualToFollowing(s, i, step)
}

// isSequenceEnd checks if the current number is the last in the sequence
func isSequenceEnd(s []int, i int, step int) bool {
	return isEqualToPrevious(s, i, step) && !isEqualToFollowing(s, i, step)
}

// isNotInSequence checks if the current number is NOT in the sequence
func isNotInSequence(s []int, i int, step int) bool {
	return !isSequenceStart(s, i, step) && !isSequenceMiddle(s, i, step) && !isSequenceEnd(s, i, step)
}
