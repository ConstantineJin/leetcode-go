package main

import "testing"

func Test_maximumBeauty(t *testing.T) {
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
				nums: []int{4, 6, 1, 2},
				k:    2,
			},
			wantAns: 3,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 1, 1, 1},
				k:    10,
			},
			wantAns: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumBeauty(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maximumBeauty() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
