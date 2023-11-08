package main

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name:    "Example1",
			args:    args{nums: []int{1, 2, 3, 4}},
			wantRes: []int{24, 12, 8, 6},
		},
		{
			name:    "Example2",
			args:    args{nums: []int{-1, 1, 0, -3, 3}},
			wantRes: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := productExceptSelf(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("productExceptSelf() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
