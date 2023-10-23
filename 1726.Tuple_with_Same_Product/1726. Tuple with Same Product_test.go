package main

import "testing"

func Test_tupleSameProduct(t *testing.T) {
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
			args:    args{nums: []int{2, 3, 4, 6}},
			wantRes: 8,
		},
		{
			name:    "Example2",
			args:    args{nums: []int{1, 2, 4, 5, 10}},
			wantRes: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := tupleSameProduct(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("tupleSameProduct() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
