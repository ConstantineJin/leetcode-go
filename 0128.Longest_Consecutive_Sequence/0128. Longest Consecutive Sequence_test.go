package main

import "testing"

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{100, 4, 200, 1, 3, 2}},
			wantRes: 4,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
			wantRes: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := longestConsecutive(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("longestConsecutive() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
