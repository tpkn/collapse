package collapse

import (
	"testing"
	
	"github.com/stretchr/testify/require"
)

func Test_DatesOnly(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{name: "Empty slice", input: []string{}, want: [][]string{}},
		{name: "1 element", input: []string{"2025-01-01"}, want: [][]string{{"2025-01-01", "2025-01-01"}}},
		{name: "2 elements / equal", input: []string{"2025-01-01", "2025-01-01"}, want: [][]string{{"2025-01-01", "2025-01-01"}}},
		{name: "2 elements / consecutive", input: []string{"2025-01-01", "2025-01-02"}, want: [][]string{{"2025-01-01", "2025-01-02"}}},
		{name: "2 elements / inconsecutive", input: []string{"2025-01-01", "2025-01-05"}, want: [][]string{{"2025-01-01", "2025-01-01"}, {"2025-01-05", "2025-01-05"}}},
		{name: "3 elements / equal", input: []string{"2025-01-01", "2025-01-01", "2025-01-01"}, want: [][]string{{"2025-01-01", "2025-01-01"}}},
		{name: "3 elements / consecutive", input: []string{"2025-01-01", "2025-01-02", "2025-01-03"}, want: [][]string{{"2025-01-01", "2025-01-03"}}},
		{name: "3 elements / inconsecutive", input: []string{"2025-01-01", "2025-01-04", "2025-01-07"}, want: [][]string{{"2025-01-01", "2025-01-01"}, {"2025-01-04", "2025-01-04"}, {"2025-01-07", "2025-01-07"}}},
		{name: "4 elements / 2 consecutive (at the beginning)", input: []string{"2025-01-01", "2025-01-02", "2025-01-05", "2025-01-07"}, want: [][]string{{"2025-01-01", "2025-01-02"}, {"2025-01-05", "2025-01-05"}, {"2025-01-07", "2025-01-07"}}},
		{name: "4 elements / 2 consecutive (in the middle)", input: []string{"2025-01-01", "2025-01-03", "2025-01-04", "2025-01-07"}, want: [][]string{{"2025-01-01", "2025-01-01"}, {"2025-01-03", "2025-01-04"}, {"2025-01-07", "2025-01-07"}}},
		{name: "4 elements / 2 consecutive (at the end)", input: []string{"2025-01-01", "2025-01-03", "2025-01-06", "2025-01-07"}, want: [][]string{{"2025-01-01", "2025-01-01"}, {"2025-01-03", "2025-01-03"}, {"2025-01-06", "2025-01-07"}}},
		{name: "4 elements / 2 consecutive and 1 is not", input: []string{"2025-01-01", "2025-01-02", "2025-01-03", "2025-01-07"}, want: [][]string{{"2025-01-01", "2025-01-03"}, {"2025-01-07", "2025-01-07"}}},
		{name: "4 elements / consecutive", input: []string{"2025-01-01", "2025-01-02", "2025-01-03", "2025-01-04"}, want: [][]string{{"2025-01-01", "2025-01-04"}}},
		{name: "Big slice", input: []string{"2025-01-01", "2025-01-02", "2025-01-03", "2025-01-10", "2025-01-11", "2025-01-14", "2025-01-17", "2025-01-20", "2025-01-21"}, want: [][]string{{"2025-01-01", "2025-01-03"}, {"2025-01-10", "2025-01-11"}, {"2025-01-14", "2025-01-14"}, {"2025-01-17", "2025-01-17"}, {"2025-01-20", "2025-01-21"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DatesOnly(tt.input)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func Benchmark_DatesOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = DatesOnly([]string{"2025-01-01", "2025-01-02", "2025-01-03", "2025-01-10", "2025-01-11", "2025-01-14", "2025-01-17", "2025-01-20", "2025-01-21"})
	}
}
