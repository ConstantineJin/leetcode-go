package main

import "testing"

func Test_countMatchingSubarrays(t *testing.T) {
	type args struct {
		nums    []int
		pattern []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				nums:    []int{1, 2, 3, 4, 5, 6},
				pattern: []int{1, 1},
			},
			wantAns: 4,
		},
		{
			name: "Example2",
			args: args{
				nums:    []int{1, 4, 4, 1, 3, 5, 5, 3},
				pattern: []int{1, 0, -1},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countMatchingSubarrays(tt.args.nums, tt.args.pattern); gotAns != tt.wantAns {
				t.Errorf("countMatchingSubarrays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
