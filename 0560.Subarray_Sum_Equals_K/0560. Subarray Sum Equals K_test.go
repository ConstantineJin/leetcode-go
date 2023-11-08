package main

import "testing"

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{1, 1, 1}, k: 2},
			wantRes: 2,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 3}, k: 3},
			wantRes: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := subarraySum(tt.args.nums, tt.args.k); gotRes != tt.wantRes {
				t.Errorf("subarraySum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
