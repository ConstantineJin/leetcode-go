package main

import "testing"

func Test_numberOfPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "Example1",
			args: args{
				nums1: []int{1, 3, 4},
				nums2: []int{1, 3, 4},
				k:     1,
			},
			wantAns: 5,
		},
		{
			name: "Example2",
			args: args{
				nums1: []int{1, 2, 4, 12},
				nums2: []int{2, 4},
				k:     3,
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfPairs(tt.args.nums1, tt.args.nums2, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("numberOfPairs() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
