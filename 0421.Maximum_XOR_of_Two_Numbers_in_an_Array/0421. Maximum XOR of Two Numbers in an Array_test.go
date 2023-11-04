package main

import "testing"

func Test_findMaximumXOR(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{3, 10, 5, 25, 2, 8}},
			wantMax: 28,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}},
			wantMax: 127,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := findMaximumXOR(tt.args.nums); gotMax != tt.wantMax {
				t.Errorf("findMaximumXOR() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
