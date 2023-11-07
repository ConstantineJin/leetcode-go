package main

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name:    "Example1",
			args:    args{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			wantRes: 6,
		},
		{
			name:    "Example2",
			args:    args{height: []int{4, 2, 0, 3, 2, 5}},
			wantRes: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := trap(tt.args.height); gotRes != tt.wantRes {
				t.Errorf("trap() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
