package collapse

// Numbers collapses a sorted slice of numbers into {{start, end}, {...}} format
func Numbers(list []int) [][]int {
	var result [][]int
	var list_len = len(list)
	var step = 1
	
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
			if isSequenceStart(list, i, step) {
				start = list[i]
				continue
			}
			if isSequenceMiddle(list, i, step) {
				continue
			}
			if isSequenceEnd(list, i, step) {
				end = list[i]
			}
			if isNotInSequence(list, i, step) {
				start = list[i]
				end = list[i]
			}
			result = append(result, []int{start, end})
		}
	}
	
	return result
}
