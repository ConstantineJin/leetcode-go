package main

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "Example1",
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			wantAns: []int{2, 2},
		},
		{
			name: "Example2",
			args: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			wantAns: []int{4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := intersect(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("intersect() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
