package main

import "testing"

func Test_numberOfSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 1, 2, 1, 1},
				k:    3,
			},
			wantAns: 2,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{2, 4, 6},
				k:    1,
			},
			wantAns: 0,
		},
		{
			name: "Example3",
			args: args{
				nums: []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2},
				k:    2,
			},
			wantAns: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numberOfSubarrays(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("numberOfSubarrays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
