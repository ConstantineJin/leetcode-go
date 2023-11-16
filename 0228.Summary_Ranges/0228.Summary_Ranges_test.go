package main

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{0, 1, 2, 4, 5, 7}},
			wantAns: []string{"0->2", "4->5", "7"},
		},
		{
			name:    "Example2",
			args:    args{nums: []int{0, 2, 3, 4, 6, 8, 9}},
			wantAns: []string{"0", "2->4", "6", "8->9"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := summaryRanges(tt.args.nums); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("summaryRanges() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
