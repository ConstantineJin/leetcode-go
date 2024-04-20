package main

import "testing"

func Test_findLength(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				nums1: []int{1, 2, 3, 2, 1},
				nums2: []int{3, 2, 1, 4, 7},
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				nums1: []int{0, 0, 0, 0, 0},
				nums2: []int{0, 0, 0, 0, 0},
			},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findLength(tt.args.nums1, tt.args.nums2); gotAns != tt.wantAns {
				t.Errorf("findLength() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
