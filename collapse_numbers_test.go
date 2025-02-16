package collapse

import (
	"testing"
	
	"github.com/stretchr/testify/require"
)

func Test_Numbers(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  [][]int
	}{
		{name: "Empty slice", input: []int{}, want: [][]int{}},
		{name: "1 element", input: []int{1}, want: [][]int{{1, 1}}},
		{name: "2 elements / equal", input: []int{1, 1}, want: [][]int{{1, 1}}},
		{name: "2 elements / consecutive", input: []int{1, 2}, want: [][]int{{1, 2}}},
		{name: "2 elements / inconsecutive", input: []int{1, 5}, want: [][]int{{1, 1}, {5, 5}}},
		{name: "3 elements / equal", input: []int{1, 1, 1}, want: [][]int{{1, 1}}},
		{name: "3 elements / consecutive", input: []int{1, 2, 3}, want: [][]int{{1, 3}}},
		{name: "3 elements / inconsecutive", input: []int{1, 5, 7}, want: [][]int{{1, 1}, {5, 5}, {7, 7}}},
		{name: "4 elements / 2 consecutive (at the beginning)", input: []int{1, 2, 5, 7}, want: [][]int{{1, 2}, {5, 5}, {7, 7}}},
		{name: "4 elements / 2 consecutive (in the middle)", input: []int{1, 3, 4, 7}, want: [][]int{{1, 1}, {3, 4}, {7, 7}}},
		{name: "4 elements / 2 consecutive (at the end)", input: []int{1, 3, 6, 7}, want: [][]int{{1, 1}, {3, 3}, {6, 7}}},
		{name: "4 elements / 2 consecutive and 1 is not", input: []int{1, 2, 3, 7}, want: [][]int{{1, 3}, {7, 7}}},
		{name: "4 elements / consecutive", input: []int{1, 2, 3, 4}, want: [][]int{{1, 4}}},
		{name: "Big slice", input: []int{0, 1, 2, 3, 10, 11, 34, 37, 90, 91}, want: [][]int{{0, 3}, {10, 11}, {34, 34}, {37, 37}, {90, 91}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Numbers(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}

func Benchmark_Numbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Numbers([]int{0, 1, 2, 3, 10, 11, 34, 37, 90, 91, 100, 112, 113, 114, 115, 116, 118, 235, 334, 335, 456, 478, 568, 777})
	}
}
