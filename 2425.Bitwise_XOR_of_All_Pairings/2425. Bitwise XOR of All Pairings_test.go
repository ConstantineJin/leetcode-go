package main

import "testing"

func Test_xorAllNums(t *testing.T) {
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
				nums1: []int{2, 1, 3},
				nums2: []int{10, 2, 5, 0},
			},
			wantAns: 13,
		},
		{
			name: "Example2",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := xorAllNums(tt.args.nums1, tt.args.nums2); gotAns != tt.wantAns {
				t.Errorf("xorAllNums() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
