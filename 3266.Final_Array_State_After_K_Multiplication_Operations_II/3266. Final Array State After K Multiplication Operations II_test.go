package main

import (
	"reflect"
	"testing"
)

func Test_getFinalState(t *testing.T) {
	type args struct {
		nums       []int
		k          int
		multiplier int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				nums:       []int{2, 1, 3, 5, 6},
				k:          5,
				multiplier: 2,
			},
			want: []int{8, 4, 6, 5, 6},
		},
		{
			name: "Example2",
			args: args{
				nums:       []int{100000, 2000},
				k:          2,
				multiplier: 1000000,
			},
			want: []int{999999307, 999999993},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFinalState(tt.args.nums, tt.args.k, tt.args.multiplier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFinalState() = %v, want %v", got, tt.want)
			}
		})
	}
}
