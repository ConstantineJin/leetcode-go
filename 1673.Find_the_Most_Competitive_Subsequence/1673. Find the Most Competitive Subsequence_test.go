package main

import (
	"reflect"
	"testing"
)

func Test_mostCompetitive(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{3, 5, 2, 6},
				k:    2,
			},
			wantAns: []int{2, 6},
		},
		{
			name: "Example2",
			args: args{
				nums: []int{2, 4, 3, 3, 5, 4, 9, 6},
				k:    4,
			},
			wantAns: []int{2, 3, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCompetitive(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.wantAns) {
				t.Errorf("mostCompetitive() = %v, want %v", got, tt.wantAns)
			}
		})
	}
}
