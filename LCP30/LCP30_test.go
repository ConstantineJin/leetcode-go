package main

import "testing"

func Test_magicTower(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{100, 100, 100, -250, -60, -140, -50, -50, 100, 150}},
			wantAns: 1,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{-200, -300, 400, 0}},
			wantAns: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := magicTower(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("magicTower() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
