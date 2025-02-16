package collapse

import (
	"time"
)

// DatesOnly collapses a sorted slice of dates ('yyyy-mm-dd') into {{start, end}, {...}} format
func DatesOnly(list []string) ([][]string, error) {
	var result [][]string
	var list_len = len(list)
	var step = 86400000 // 1 day in ms
	
	// Transform date to unix format
	var list_ms []int
	for _, d := range list {
		parsed, err := time.Parse(time.DateOnly, d)
		if err != nil {
			return nil, err
		}
		list_ms = append(list_ms, int(parsed.UnixMilli()))
	}
	
	switch list_len {
	case 0:
		result = [][]string{}
	case 1:
		result = [][]string{{list[0], list[0]}}
	case 2:
		if list_ms[0] == list_ms[1] || list_ms[0] == list_ms[1]-step {
			result = [][]string{{list[0], list[1]}}
		} else {
			result = [][]string{{list[0], list[0]}, {list[1], list[1]}}
		}
	default:
		var start, end string
		for i := 0; i < list_len; i++ {
			if isSequenceStart(list_ms, i, step) {
				start = list[i]
				continue
			}
			if isSequenceMiddle(list_ms, i, step) {
				continue
			}
			if isSequenceEnd(list_ms, i, step) {
				end = list[i]
			}
			if isNotInSequence(list_ms, i, step) {
				start = list[i]
				end = list[i]
			}
			result = append(result, []string{start, end})
		}
	}
	
	return result, nil
}
