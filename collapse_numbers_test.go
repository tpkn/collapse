package collapse

import (
	"reflect"
	"testing"
)

func Test_Numbers(t *testing.T) {
	type args struct {
		ip_list []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "Empty slice", args: args{ip_list: []int{}}, want: [][]int{}},
		{name: "1 element", args: args{ip_list: []int{1}}, want: [][]int{{1, 1}}},
		{name: "2 elements / equal", args: args{ip_list: []int{1, 1}}, want: [][]int{{1, 1}}},
		{name: "2 elements / consecutive", args: args{ip_list: []int{1, 2}}, want: [][]int{{1, 2}}},
		{name: "2 elements / inconsecutive", args: args{ip_list: []int{1, 5}}, want: [][]int{{1, 1}, {5, 5}}},
		{name: "3 elements / equal", args: args{ip_list: []int{1, 1, 1}}, want: [][]int{{1, 1}}},
		{name: "3 elements / consecutive", args: args{ip_list: []int{1, 2, 3}}, want: [][]int{{1, 3}}},
		{name: "3 elements / inconsecutive", args: args{ip_list: []int{1, 5, 7}}, want: [][]int{{1, 1}, {5, 5}, {7, 7}}},
		{name: "4 elements / 2 consecutive (at the beginning)", args: args{ip_list: []int{1, 2, 5, 7}}, want: [][]int{{1, 2}, {5, 5}, {7, 7}}},
		{name: "4 elements / 2 consecutive (in the middle)", args: args{ip_list: []int{1, 3, 4, 7}}, want: [][]int{{1, 1}, {3, 4}, {7, 7}}},
		{name: "4 elements / 2 consecutive (at the end)", args: args{ip_list: []int{1, 3, 6, 7}}, want: [][]int{{1, 1}, {3, 3}, {6, 7}}},
		{name: "4 elements / 2 consecutive and 1 is not", args: args{ip_list: []int{1, 2, 3, 7}}, want: [][]int{{1, 3}, {7, 7}}},
		{name: "4 elements / consecutive", args: args{ip_list: []int{1, 2, 3, 4}}, want: [][]int{{1, 4}}},
		{name: "Big slice", args: args{ip_list: []int{0, 1, 2, 3, 10, 11, 34, 37, 90, 91}}, want: [][]int{{0, 3}, {10, 11}, {34, 34}, {37, 37}, {90, 91}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Numbers(tt.args.ip_list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Numbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_Numbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Numbers([]int{0, 1, 2, 3, 10, 11, 34, 37, 90, 91, 100, 112, 113, 114, 115, 116, 118, 235, 334, 335, 456, 478, 568, 777})
	}
}
