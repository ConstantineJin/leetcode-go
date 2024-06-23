package main

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				nums:  []int{8, 2, 4, 7},
				limit: 4,
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				nums:  []int{10, 1, 2, 4, 7, 2},
				limit: 5,
			},
			wantAns: 4,
		},
		{
			name: "Example3",
			args: args{
				nums:  []int{4, 2, 2, 2, 4, 4, 2, 2},
				limit: 0,
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestSubarray(tt.args.nums, tt.args.limit); gotAns != tt.wantAns {
				t.Errorf("longestSubarray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
